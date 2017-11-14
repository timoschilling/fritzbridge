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

  ticker := time.NewTicker(time.Millisecond * 1000)
  go func() {
    for _ = range ticker.C {
      for _, device := range api.GetDevices(config).Device {
        accessory, err := FindDevice(device.Identifier, thermostats)
        if err != false {
          accessory.Thermostat.CurrentTemperature.SetValue(device.GetCurrentTemperature())
          accessory.Thermostat.TargetTemperature.SetValue(device.GetTargetTemperature())
          accessory.Thermostat.CurrentHeatingCoolingState.SetValue(device.GetCurrentHeatingCoolingState())
          accessory.Thermostat.TargetHeatingCoolingState.SetValue(device.GetCurrentHeatingCoolingState())
        }
      }
    }
  }()

  hc.OnTermination(func() {
    <-t.Stop()
  })

  t.Start()
}
