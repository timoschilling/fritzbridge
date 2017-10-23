package api

import (
  "github.com/philippfranke/go-fritzbox/fritzbox"
  "fmt"
  "log"
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
