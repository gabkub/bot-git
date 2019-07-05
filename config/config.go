package config

import (
	"encoding/json"
	"io/ioutil"
)

type BotConfig struct {
	Name       string `json:"Name"`
	Password   string `json:"Password"`
	Email      string `json:"Email"`
	TeamName   string `json:"TeamName"`
	EnglishDay string `json:"EnglishDay"`
}

func Read(path string) (*BotConfig, error) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}
	cfg := &BotConfig{}
	e = json.Unmarshal([]byte(file), cfg)
	if e != nil {
		return nil, e
	}
	return cfg, nil
}