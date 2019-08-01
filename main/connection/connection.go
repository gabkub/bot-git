package connection

import (
	"bot-git/bot/limit"
	"bot-git/config"
	"bot-git/logg"
	"fmt"
	"github.com/mattermost/mattermost-server/model"
	"log"
	"strings"
)

var Websocket *model.WebSocketClient
var protocol = "http"
var secure = false

func Connect() {

	config.BotCfg.Port = strings.ToLower(config.BotCfg.Port)
	if config.BotCfg.Port == "443" {
		protocol = "https"
		secure = true
	}

	connectServer()

	loginAsTheBotUser()
	setBotTeam()

	if limit.Users == nil {
		limit.SetUsersList()
	}

	connectWebsocket()
	Websocket.Listen()
}

func connectServer() {
		config.ConnectionCfg.Client = model.NewAPIv4Client(fmt.Sprintf("%s://%s:%s", protocol, config.BotCfg.Server, config.BotCfg.Port))
		if config.ConnectionCfg.Client == nil {
			logg.WriteToFile(fmt.Sprintf("Error while connecting to the Mattermost API. Connecting again."))
			log.Println(fmt.Sprintf("Error while connecting to the Mattermost API. Connecting again."))
		}
		makeSureServerIsRunning()
}

func makeSureServerIsRunning() {

	for {
		if _, resp := config.ConnectionCfg.Client.GetPing(); resp.Error != nil {
			logg.WriteToFile(fmt.Sprintf("Error pinging the Mattermost server %s. Details: %s", config.ConnectionCfg.Client.Url, resp.Error.Message))
			log.Println(fmt.Sprintf("Error pinging the Mattermost server %s. Details: %s", config.ConnectionCfg.Client.Url, resp.Error.Message))
		} else {
			logg.WriteToFile(fmt.Sprintf("Mattermost server %s pinged successfully.", config.ConnectionCfg.Client.Url))
			break
		}
	}
}

func loginAsTheBotUser() {
	if 	user,resp := config.ConnectionCfg.Client.Login(config.BotCfg.BotName, config.BotCfg.Password); resp.Error != nil {
		logg.WriteToFile("There was a problem logging into the Mattermost server. Details: " + resp.Error.Message)
		log.Fatal("There was a problem logging into the Mattermost server. Details: " + resp.Error.Message)
	} else {
		logg.WriteToFile("Bot logged into the Mattermost server successfully.")
		config.ConnectionCfg.BotUser = user
	}

	revokePreviousSessions()
}

func revokePreviousSessions() {

	if sessions,_ := config.ConnectionCfg.Client.GetSessions(config.ConnectionCfg.BotUser.Id,""); sessions != nil {
		for i,session := range sessions {
			if i != 0 {
				config.ConnectionCfg.Client.RevokeSession(config.ConnectionCfg.BotUser.Id, session.Id)
			}
		}
	}
}

func setBotTeam() {
	if team, resp := config.ConnectionCfg.Client.GetTeamByName(config.BotCfg.TeamName,""); resp.Error != nil {
		logg.WriteToFile(fmt.Sprintf("Team '%s' does not exist.",config.BotCfg.TeamName))
		log.Fatal(fmt.Sprintf("Team '%s' does not exist.",config.BotCfg.TeamName))
	} else {
		config.ConnectionCfg.BotTeam = team
	}
}

func connectWebsocket() {

	ws := "ws"
	if secure {
		ws = "wss"
	}

	for {
		websocket, err := model.NewWebSocketClient4(fmt.Sprintf("%s://%s:%s", ws, config.BotCfg.Server, config.BotCfg.Port), config.ConnectionCfg.Client.AuthToken)
		if err != nil {
			logg.WriteToFile("Error connecting to the web socket. Details: " + err.DetailedError)
		} else {
			Websocket = websocket
			break
		}
	}
}