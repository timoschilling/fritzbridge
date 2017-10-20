package main

import (
  "github.com/brutella/hc"
  "github.com/brutella/hc/accessory"
  "log"
)

func main() {
  info := accessory.Info{
    Name:         "Fritz!Bridge",
    Manufacturer: "Timo Schilling",
  }

  acc := accessory.NewThermostat(info, 20, 16, 28, 0.5)

  info2 := accessory.Info{
    Name:         "Room 2",
    Manufacturer: "Timo Schilling",
  }

  acc2 := accessory.NewThermostat(info2, 20, 16, 28, 0.5)

  t, err := hc.NewIPTransport(hc.Config{Pin: "12341234", StoragePath: "database"}, acc.Accessory, acc2.Accessory)
  if err != nil {
    log.Fatal(err)
  }

  hc.OnTermination(func() {
    t.Stop()
  })

  t.Start()
}
