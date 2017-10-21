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

    c := 19 + i
    thermostat := accessory.NewThermostat(info, float64(c), 16, 28, 0.5)
    state := 0
    if c < 20 {
      state = 1
    }
    state = 1
    thermostat.Thermostat.TargetTemperature.SetValue(float64(c + 3))
    thermostat.Thermostat.CurrentHeatingCoolingState.SetValue(state)

    accessories = append(accessories, thermostat.Accessory)
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
