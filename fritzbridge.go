package main

import (
  "github.com/brutella/hc"
  "github.com/brutella/hc/accessory"
  "log"
  "time"
  "./api"
  "./bridge"
)

func main() {
  config := api.GetConfig()
  thermostats := []*accessory.Thermostat{}

  for _, device := range api.GetDevices(config).Device {
    info := accessory.Info{
      Name:         device.Name,
      Manufacturer: device.Manufacturer,
      SerialNumber: device.Identifier,
      Model:        device.Productname,
    }

    thermostat := accessory.NewThermostat(info, device.GetCurrentTemperature(), 16, 28, 0.5)
    thermostat.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(target_temperature float64){
      api.SetTargetTemperature(device.Identifier, target_temperature, config)
    })

    thermostats = append(thermostats, thermostat)
  }

  accessories := []*accessory.Accessory{}
  for _, thermostat := range thermostats {
    accessories = append(accessories, thermostat.Accessory)
  }

  t, err := hc.NewIPTransport(hc.Config{Pin: config.Pin, StoragePath: "database"}, bridge.NewBridge().Accessory, accessories...)
  if err != nil {
    log.Fatal(err)
  }

  ticker := time.NewTicker(time.Millisecond * 1000)
  go func() {
    for _ = range ticker.C {
      for i, device := range api.GetDevices(config).Device {
        accessory := thermostats[i]
        accessory.Thermostat.CurrentTemperature.SetValue(device.GetCurrentTemperature())
        accessory.Thermostat.TargetTemperature.SetValue(device.GetTargetTemperature())
        accessory.Thermostat.CurrentHeatingCoolingState.SetValue(device.GetCurrentHeatingCoolingState())
        accessory.Thermostat.TargetHeatingCoolingState.SetValue(device.GetCurrentHeatingCoolingState())
      }
    }
  }()

  hc.OnTermination(func() {
    <-t.Stop()
  })

  t.Start()
}
