package config

import (
	"encoding/json"
	"github.com/mattermost/mattermost-server/model"
	"io/ioutil"
)
type MMConfig struct{
	Client           *model.Client4
	WebSocketClient  *model.WebSocketClient
	BotUser          *model.User
	BotTeam          *model.Team
	BotConfig		 *model.Config
}

type BotConfig struct {
	Server	   string `json:"Server"`
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