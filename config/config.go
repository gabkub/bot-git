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
	WebSocketClient  *model.WebSocketClient
	BotUser          *model.User
	BotTeam          *model.Team
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
	Text   string
	Img    Image
	IsJoke bool
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

func Read() BotConfig {

	var path string
	if len(os.Args) < 2 {
		path = "./config.json"
	} else {
		path = os.Args[1]
	}

	file, e := ioutil.ReadFile(path)
	if e != nil {
		log.Fatal("Error while opening the configuration file.")
		os.Exit(1)
	}
	cfg := &BotConfig{}
	e = json.Unmarshal([]byte(file), cfg)
	if e != nil {
		log.Fatal("Error while reading the configuration file.")
		os.Exit(1)
	}
	return *cfg
}