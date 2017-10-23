package api

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
  Celsius         float64  `xml:"temperature>celsius"`
  Offset          float64  `xml:"temperature>offset"`

  // hkr
  Tist            string   `xml:"hkr>tist"`
  Tsoll           string   `xml:"hkr>tsoll"`
  Absenk          string   `xml:"hkr>absenk"`
  Komfort         string   `xml:"hkr>komfort"`
  Lock            string   `xml:"hkr>lock"`
  Devicelock      string   `xml:"hkr>devicelock"`
  Errorcode       string   `xml:"hkr>errorcode"`
  Batterylow      string   `xml:"hkr>batterylow"`

  // nextchange
  Endperiod       string   `xml:"hkr>nextchange>endperiod"`
  Tchange         string   `xml:"hkr>nextchange>tchange"`
}

func (d *Device) GetCelsius() float64 {
  return d.Celsius / 10.0
}
