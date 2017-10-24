package main

import (
  "github.com/brutella/hc"
  "github.com/brutella/hc/accessory"
  "log"
  "./api"
  "./bridge"
)

func main() {
  config := api.GetConfig()
  accessories := []*accessory.Accessory{}

  devices := api.GetDevices(config)

  for i := 0; i < len(devices.Device); i++ {
    device := devices.Device[i]
    info := accessory.Info{
      Name:         device.Name,
      Manufacturer: device.Manufacturer,
      SerialNumber: device.Identifier,
      Model:        device.Productname,
    }

    thermostat := accessory.NewThermostat(info, device.GetCurrentTemperature(), 16, 28, 0.5)

    accessories = append(accessories, thermostat.Accessory)
  }

  t, err := hc.NewIPTransport(hc.Config{Pin: config.Pin, StoragePath: "database"}, bridge.NewBridge().Accessory, accessories...)
  if err != nil {
    log.Fatal(err)
  }

  hc.OnTermination(func() {
    <-t.Stop()
  })

  t.Start()
}
