package api

import (
  "math"
  "github.com/brutella/hc/characteristic"
)

type Devicelist struct {
  Version         string   `xml:"version,attr"`
  Device          []Device `xml:"device"`
}

type Device struct {
  Identifier      string   `xml:"identifier,attr"`
  Id              int      `xml:"id,attr"`
  Functionbitmask int      `xml:"functionbitmask,attr"`
  Fwversion       string   `xml:"fwversion,attr"`
  Manufacturer    string   `xml:"manufacturer,attr"`
  Productname     string   `xml:"productname,attr"`

  Present         bool     `xml:"present"`
  Name            string   `xml:"name"`

  // temperature
  Celsius         int      `xml:"temperature>celsius"`
  Offset          int      `xml:"temperature>offset"`

  // hkr
  Tist            int      `xml:"hkr>tist"`
  Tsoll           int      `xml:"hkr>tsoll"`
  Absenk          int      `xml:"hkr>absenk"`
  Komfort         int      `xml:"hkr>komfort"`
  Lock            bool     `xml:"hkr>lock"`
  Devicelock      bool     `xml:"hkr>devicelock"`
  Errorcode       int      `xml:"hkr>errorcode"`
  Batterylow      bool     `xml:"hkr>batterylow"`

  // nextchange
  Endperiod       int      `xml:"hkr>nextchange>endperiod"`
  Tchange         int      `xml:"hkr>nextchange>tchange"`
}

func (d *Device) GetCurrentTemperature() float64 {
  return math.Floor(float64(d.Celsius) / 10.0)
}

func (d *Device) GetTargetTemperature() float64 {
  if d.Tsoll == 253 {
    return d.GetCurrentTemperature()
  }
  if d.Tsoll == 254 {
    28.0
  }
  return math.Floor(float64(d.Tsoll) / 2.0)
}

func (d *Device) GetTargetHeatingCoolingState() int {
  if d.Tsoll == 253 {
    return characteristic.TargetHeatingCoolingStateOff
  } else if d.GetCurrentTemperature() < d.GetTargetTemperature() {
    return characteristic.TargetHeatingCoolingStateHeat
  } else {
    return characteristic.TargetHeatingCoolingStateAuto
  }
}

func (d *Device) GetStatusLowBattery() int {
  if d.Batterylow {
    return characteristic.StatusLowBatteryBatteryLevelLow
  } else {
    return characteristic.StatusLowBatteryBatteryLevelNormal
  }
}

func (d *Device) GetBatteryLevel() int {
  if d.Batterylow {
    return 10
  } else {
    return 80
  }
}
