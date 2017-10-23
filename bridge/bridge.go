package bridge

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
    Model:        "Fritz!Box",
  }
  acc.Accessory = accessory.New(info, accessory.TypeBridge)

  return &acc
}
