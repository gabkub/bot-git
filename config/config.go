package config

import (
	"encoding/json"
	"../meme"
	"github.com/mattermost/mattermost-server/model"
	"io/ioutil"
)

// Mattermost connection data
var MmCfg MMConfig

// bot user data
var BotCfg BotConfig

type MMConfig struct{
	Client           *model.Client4
	WebSocketClient  *model.WebSocketClient
	BotUser          *model.User
	BotTeam          *model.Team
	BotConfig		 *model.Config
}

type BotConfig struct {
	Server	   string `json:"Server:port"`
	Name       string `json:"Name"`
	Password   string `json:"Password"`
	Email      string `json:"Email"`
	TeamName   string `json:"TeamName"`
	EnglishDay string `json:"EnglishDay"`
}

type Msg struct {
	Text string
	Img meme.Meme
}

func Read(path string) BotConfig {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return BotConfig{}
	}
	cfg := &BotConfig{}
	e = json.Unmarshal([]byte(file), cfg)
	if e != nil {
		return BotConfig{}
	}
	return *cfg
}