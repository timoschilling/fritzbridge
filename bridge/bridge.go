package bridge

import (
  "github.com/brutella/hc/accessory"
)

type Bridge struct {
  *accessory.Accessory
}

func NewBridge(name string) *Bridge {
  acc := Bridge{}
  info := accessory.Info{
    Name:         name,
    Manufacturer: "AVM",
    Model:        "Fritz!Box",
  }
  acc.Accessory = accessory.New(info, accessory.TypeBridge)

  return &acc
}
