package config

import (
	"encoding/json"
	"github.com/mattermost/mattermost-server/model"
	"io/ioutil"
	"log"
	"os"
)

// Mattermost connection data
var MmCfg MMConfig

// bot user data
var BotCfg = Read()

type MMConfig struct{
	Client           *model.Client4
	Team 			 *model.Team
	WebSocketClient  *model.WebSocketClient
	BotUser          *model.User
	BotTeam          *model.Team
	Session			 *model.Session
}

type BotConfig struct {
	Server	   string `json:"Server"`
	Port   	   string `json:"Port"`
	BotName    string `json:"BotName"`
	Password   string `json:"Password"`
	TeamName   string `json:"TeamName"`
	EnglishDay string `json:"EnglishDay"`
}
func Read() BotConfig {

	var path string
	if len(os.Args) < 2 {
		path = "./config.json"
	} else {
		if os.Args[1]=="-test.v" {
			path = "./config.json"
		} else {
			path = os.Args[1]
		}
	}

	file, e := ioutil.ReadFile(path)
	if e != nil {
		log.Fatal("Error while opening the configuration file. Path: "+path)
	}
	cfg := &BotConfig{}
	e = json.Unmarshal([]byte(file), cfg)
	if e != nil {
		log.Fatal("Error while reading the configuration file.")
	}
	return *cfg
}

