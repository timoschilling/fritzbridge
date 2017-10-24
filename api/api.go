package api

import (
  "github.com/philippfranke/go-fritzbox/fritzbox"
  "fmt"
  "log"
  "strings"
  "net/url"
  "net/http"
  "io/ioutil"
  "encoding/xml"
)

func GetSessionId(config *Config) string {
  c := fritzbox.NewClient(nil)
  c.BaseURL, _ = url.Parse(config.Url)
  err := c.Auth(config.Username, config.Password)

  if err != nil {
    log.Fatalf("Auth failed: %v", err)
  }

  return c.String()
}

func GetDevices(config *Config) *Devicelist {
  resp, err := http.Get(fmt.Sprintf("%s/webservices/homeautoswitch.lua?switchcmd=getdevicelistinfos&sid=%s", config.Url, GetSessionId(config)))
  if err != nil {
    log.Fatalf("Request failed: %v", err)
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalf("Read body failed: %v", err)
  }

  devices := &Devicelist{}
  xml.Unmarshal(body, &devices)

  return devices
}

func SetTargetTemperature(ain string, target_temperature float64, config *Config)  {
  tist := int(target_temperature * 2)
  ain = strings.Replace(ain, " ", "%20", -1)
  log.Println(ain)
  url := fmt.Sprintf("%s/webservices/homeautoswitch.lua?switchcmd=sethkrtsoll&ain=%s&param=%d&sid=%s", config.Url, ain, tist, GetSessionId(config))
  log.Println(url)
  _, err := http.Get(url)
  if err != nil {
    log.Fatalf("Request failed: %v", err)
  }
}
