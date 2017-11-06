package api

import (
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
  return float64(d.Celsius) / 10.0
}

func (d *Device) GetTargetTemperature() float64 {
  return float64(d.Tsoll) / 2.0
}

func (d *Device) GetCurrentHeatingCoolingState() int {
  if d.GetCurrentTemperature() < d.GetTargetTemperature() {
    return characteristic.CurrentHeatingCoolingStateHeat
  } else {
    return characteristic.CurrentHeatingCoolingStateOff
  }
}
