package api

import (
  "log"
  "encoding/json"
  "io/ioutil"
)

type Config struct {
  Url      string `json:"url"`
  Username string `json:"username"`
  Password string `json:"password"`
}

func GetConfig() *Config {
  data, err := ioutil.ReadFile("fritzbridge.json")
  if err != nil {
    log.Fatalf("Config read failed: %v", err)
  }

  config := &Config{}
  err = json.Unmarshal(data, &config)

  if err != nil {
    log.Fatalf("Config load failed: %v", err)
  }

  return config
}
