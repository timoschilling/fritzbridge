package main

import (
  "github.com/brutella/hc"
  "github.com/brutella/hc/accessory"
  "log"
)

func main() {
  var accessories [2]*accessory.Accessory

  info := accessory.Info{
    Name:         "Fritz!Bridge",
    Manufacturer: "Timo Schilling",
  }

  accessories[0] = accessory.NewThermostat(info, 20, 16, 28, 0.5).Accessory

  info2 := accessory.Info{
    Name:         "Room 2",
    Manufacturer: "Timo Schilling",
  }

  accessories[1] = accessory.NewThermostat(info2, 20, 16, 28, 0.5).Accessory

  t, err := hc.NewIPTransport(hc.Config{Pin: "12341234", StoragePath: "database"}, accessories[0], accessories[1])
  if err != nil {
    log.Fatal(err)
  }

  hc.OnTermination(func() {
    t.Stop()
  })

  t.Start()
}
