package connection

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-server/model"
	"log"
	"strings"
)
var Websocket *model.WebSocketClient
var secure = false
// connect with the Mattermost server
func ConnectServer() {
	protocol := "http"

	config.BotCfg.Port = strings.ToLower(config.BotCfg.Port)

	if config.BotCfg.Port == "443" {
		protocol = "https"
		secure = true
	}

	config.MmCfg.Client = model.NewAPIv4Client(fmt.Sprintf("%s://%s:%s", protocol, config.BotCfg.Server, config.BotCfg.Port))
	if config.MmCfg.Client == nil {

	}

	// test to see if the mattermost server is up and running
	makeSureServerIsRunning()

	// attempt to login to the Mattermost server as the bot user
	// This will set the token required for all future calls
	// You can get this token with client.AuthToken
	loginAsTheBotUser()
	revokePreviousSessions()

	// find the bot team
	findBotTeam()

	limit.SetTeamMembers()
	// create new WebSocket client

	ConnectWebsocket()
}

func ConnectWebsocket() {
	if Websocket != nil {
		Websocket.Close()
	}

	var err *model.AppError
	ws := "ws"
	if secure {
		ws = "wss"
	}

	Websocket, err = model.NewWebSocketClient4(fmt.Sprintf("%s://%s:%s", ws, config.BotCfg.Server, config.BotCfg.Port), config.MmCfg.Client.AuthToken)
	if err != nil {
		log.Fatal("Error connecting to the web socket. Details: " + err.DetailedError)
	}

	Websocket.Listen()
}
// check the mattermost server
func makeSureServerIsRunning() {
	if _, resp := config.MmCfg.Client .GetPing(); resp.Error != nil {
		log.Fatal(fmt.Sprintf("Error pinging the Mattermost server %s. Details: %s", config.MmCfg.Client.Url, resp.Error.Message))
	} else {
		log.Println(fmt.Sprintf("Mattermost server %s detected and running ver. %s.", config.MmCfg.Client.Url, resp.ServerVersion))
	}
}

func revokePreviousSessions() {

	if sessions,_ := config.MmCfg.Client.GetSessions(config.MmCfg.BotUser.Id,""); sessions != nil {
		for i,session := range sessions {
			if i != 0 {
				config.MmCfg.Client.RevokeSession(config.MmCfg.BotUser.Id, session.Id)
			}
		}
	}
}

// login to the chat as bot user using credentials from config
func loginAsTheBotUser() {
	if 	user,resp := config.MmCfg.Client.Login(config.BotCfg.BotName, config.BotCfg.Password); resp.Error != nil {
		log.Fatal("There was a problem logging into the Mattermost server. Details: " + resp.Error.Message)
	} else {
		//
		log.Println("Bot logged into the Mattermost server successfully.")
		config.MmCfg.BotUser = user
	}
}

// check whether the bot user is a member of the team specified in config
func findBotTeam() {
	if team, resp := config.MmCfg.Client.GetTeamByName(config.BotCfg.TeamName,""); resp.Error != nil {
		log.Fatal(fmt.Sprintf("Team '%s' does not exist.",config.BotCfg.TeamName))
	} else {
		config.MmCfg.BotTeam = team
	}
}
