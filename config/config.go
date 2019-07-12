package config

import (
	"encoding/json"
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
	Server	   string `json:"Server"`
	Protocol   string `json:"Protocol"`
	Name       string `json:"Name"`
	Password   string `json:"Password"`
	TeamName   string `json:"TeamName"`
	EnglishDay string `json:"EnglishDay"`
}

type Msg struct {
	Text string
	Img Image
}

type Image struct{
	Header 		string
	ImageUrl 	string
}

func (i Image) IsEmpty() bool{

	if i.Header == "" && i.ImageUrl == ""{
		return true
	}
	return false
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