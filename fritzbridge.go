package main

import (
  "github.com/brutella/hc"
  "github.com/brutella/hc/accessory"
  "log"
  "fmt"
)

func main() {
  accessories := []*accessory.Accessory{}

  for i := 0; i < 3; i++ {
    info := accessory.Info{
      Name:         fmt.Sprintf("Room %d", i + 1),
      Manufacturer: "Timo Schilling",
    }

    accessories = append(accessories, accessory.NewThermostat(info, 20, 16, 28, 0.5).Accessory)
  }

  t, err := hc.NewIPTransport(hc.Config{Pin: "12341234", StoragePath: "database"}, accessories[0], accessories[1:len(accessories)]...)
  if err != nil {
    log.Fatal(err)
  }

  hc.OnTermination(func() {
    t.Stop()
  })

  t.Start()
}
