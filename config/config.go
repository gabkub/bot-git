package config

import (
	"encoding/json"
	"github.com/mattermost/mattermost-server/model"
	"io/ioutil"
	"log"
	"os"
)

var ConnectionCfg connectionConfig
var BotCfg *BotConfig
var DbCfg *DbConfig

type connectionConfig struct{
	Client           *model.Client4
	Team 			 *model.Team
	WebSocketClient  *model.WebSocketClient
	BotUser          *model.User
	BotTeam          *model.Team
	Session			 *model.Session
}

type Data struct {
	BotConfig *BotConfig `json:"BotConfig"`
	DbConfig *DbConfig   `json:"DbConfig"`
}

type BotConfig struct {
	Server	   string `json:"Server"`
	Port   	   string `json:"Port"`
	BotName    string `json:"BotName"`
	Password   string `json:"Password"`
	TeamName   string `json:"TeamName"`
	EnglishDay string `json:"EnglishDay"`
}

type DbConfig struct {
	Name                 string `json:"Name"`
	Server               string `json:"Server"`
	Port                 int    `json:"Port"`
	User                 string `json:"User"`
	Password             string `json:"Password"`
	ConnectionsWarning   int    `json:"Connections_warning"`
	ConnectionsCheckCron string `json:"Connections_check_cron"`
	ConnectionsLogCron   string `json:"Connections_log_cron"`
}

func ReadConfig(aesKey string) {
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

	cfg := &Data{}
	e = json.Unmarshal([]byte(file), cfg)
	if e != nil {
		log.Fatal("Error while reading the configuration file.")
	}

	BotCfg = cfg.BotConfig
	DbCfg = cfg.DbConfig
}