package main

import (
  "os"
  "log"
  "time"
  "flag"
  "./api"
  "./bridge"
  "github.com/brutella/hc"
  "github.com/brutella/hc/accessory"
  "github.com/brutella/hc/service"
  "github.com/brutella/hc/characteristic"
)

func FindDevice(identifier string, thermostats []*accessory.Thermostat) (*accessory.Thermostat, bool) {
  for _, thermostat := range thermostats {
    if thermostat.Info.SerialNumber.GetValue() == identifier {
      return thermostat, true
    }
  }
  return nil, false
}

func main() {
  get_session_id := flag.Bool("sessionid", false, "get a sessionid")
  flag.Parse()

  config := api.GetConfig()
  config.SessionId = api.GetSessionId(config)

  if *get_session_id {
    log.Println(config.SessionId)
    os.Exit(0)
  }

  thermostats := []*accessory.Thermostat{}

  for _, device := range api.GetDevices(config).Device {
    info := accessory.Info{
      Name:         device.Name,
      Manufacturer: device.Manufacturer,
      SerialNumber: device.Identifier,
      Model:        device.Productname,
    }

    thermostat := accessory.NewThermostat(info, device.GetCurrentTemperature(), 8, 28, 0.5)
    thermostat.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(target_temperature float64){
      api.SetTargetTemperature(thermostat.Info.SerialNumber.GetValue(), target_temperature, config)
    })

    thermostat.Info.FirmwareRevision.SetValue(device.Fwversion)

    battery := service.NewBatteryService()
    thermostat.AddService(battery.Service)

    thermostats = append(thermostats, thermostat)
  }

  accessories := []*accessory.Accessory{}
  for _, thermostat := range thermostats {
    accessories = append(accessories, thermostat.Accessory)
  }

  t, err := hc.NewIPTransport(hc.Config{Pin: config.Pin, StoragePath: "database"}, bridge.NewBridge(config.BridgeName).Accessory, accessories...)
  if err != nil {
    log.Fatal(err)
  }

  ticker := time.NewTicker(time.Second * 10)
  go func() {
    for _ = range ticker.C {
      for _, device := range api.GetDevices(config).Device {
        accessory, err := FindDevice(device.Identifier, thermostats)
        if err != false {
          accessory.Thermostat.CurrentTemperature.SetValue(device.GetCurrentTemperature())
          accessory.Thermostat.TargetTemperature.SetValue(device.GetTargetTemperature())
          accessory.Thermostat.CurrentHeatingCoolingState.SetValue(device.GetCurrentHeatingCoolingState())
          accessory.Thermostat.TargetHeatingCoolingState.SetValue(device.GetTargetHeatingCoolingState())
          for _, s := range accessory.Services {
            if s.Type == service.TypeBatteryService {
              for _, c := range s.Characteristics {
                switch c.Type {
                case characteristic.TypeStatusLowBattery:
                  c.UpdateValue(device.GetStatusLowBattery())
                case characteristic.TypeBatteryLevel:
                  c.UpdateValue(device.GetBatteryLevel())
                }
              }
            }
          }
        }
      }
    }
  }()

  hc.OnTermination(func() {
    <-t.Stop()
  })

  t.Start()
}
