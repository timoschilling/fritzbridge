package main

import (
  "github.com/brutella/hc/accessory"
)

type Bridge struct {
  *accessory.Accessory
}

func NewBridge() *Bridge {
  acc := Bridge{}
  info := accessory.Info{
    Name:         "Fritz!Bridge",
    Manufacturer: "AVM",
    Model:        "Fritz!Box 7490",
  }
  acc.Accessory = accessory.New(info, accessory.TypeBridge)

  return &acc
}
