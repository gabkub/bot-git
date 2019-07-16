package main

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-server/model"
	"log"
	"os"
	"strings"
)

const VER = "1.0.2.0"
// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client
func main() {
	log.Printf("Running bot v.%v...\n", commands.VER)

	// WebSocket initialization
	websocket := connection()

	// start listening on all channels
	websocket.Listen()

	bot.Start(websocket)

}

// print details of an error
//func PrintError(err *model.AppError) {
//	println("\tError Details:")
//	println("\t\t" + err.Message)
//	println("\t\t" + err.Id)
//	println("\t\t" + err.DetailedError)
//}

// connect with the Mattermost server
func connection() *model.WebSocketClient{

	var port string
	secure := false

	config.BotCfg.Protocol = strings.ToLower(config.BotCfg.Protocol)
	switch config.BotCfg.Protocol {
	case "http":
		port = "80"
	case "https":
		port = "443"
		secure = true
	case "mattermost":
		port = "8065"
		config.BotCfg.Protocol = "http"
	}

	config.MmCfg.Client = model.NewAPIv4Client(fmt.Sprintf("%s://%s:%s", strings.ToLower(config.BotCfg.Protocol), config.BotCfg.Server, port))

	// test to see if the mattermost server is up and running

	checkProtocol()
	makeSureServerIsRunning()
	// attempt to login to the Mattermost server as the bot user
	// This will set the token required for all future calls
	// You can get this token with client.AuthToken
	loginAsTheBotUser()

	// find the bot team
	findBotTeam()
	limit.SetTeamMembers()
	// create new WebSocket client
	var err *model.AppError
	ws := "ws"
	if secure {
		ws = "wss"
	}
	config.MmCfg.WebSocketClient, err = model.NewWebSocketClient4(fmt.Sprintf("%s://%s:%s", ws, config.BotCfg.Server, port), config.MmCfg.Client.AuthToken)
	if err != nil {
		log.Fatal("Error connecting to the web socket. Details: " + err.DetailedError)
	}

	return config.MmCfg.WebSocketClient
}

func checkProtocol() {
	if config.BotCfg.Protocol != "http" && config.BotCfg.Protocol != "https" && config.BotCfg.Protocol != "mattermost" {
		log.Fatal("Protocol is not HTTP/HTTPS or Mattermost default protocol.")
		os.Exit(1)
	}
}
// check the mattermost server
func makeSureServerIsRunning() {
	if props, resp := config.MmCfg.Client.GetOldClientConfig(""); resp.Error != nil {
		log.Fatal("Error logging into the Mattermost server. Details: " + resp.Error.DetailedError)
		os.Exit(1)
	} else {
		log.Println("Mattermost server detected and is running version " + props["Version"])
	}
}

// login to the chat as bot user using credentials from config
func loginAsTheBotUser() {
	if user, resp := config.MmCfg.Client.Login(config.BotCfg.Name, config.BotCfg.Password); resp.Error != nil {
		log.Fatal("There was a problem logging into the Mattermost server. Details: "+resp.Error.DetailedError)
		os.Exit(1)
	} else {
		log.Println("Bot logged into the Mattermost server successfully.")
		config.MmCfg.BotUser = user
	}
}

// check whether the bot user is a member of the team specified in config
func findBotTeam() {
	if team, resp := config.MmCfg.Client.GetTeamByName(config.BotCfg.TeamName, ""); resp.Error != nil {
		log.Fatal("Bot appears not to be a member of the team '" + config.BotCfg.TeamName + "'")
		os.Exit(1)
	} else {
		config.MmCfg.BotTeam = team
	}
}