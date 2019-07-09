package main

import (
	"../bot"
	"../config"
	"github.com/mattermost/mattermost-server/model"
	"os"
)

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client
func main() {

	// read bot configuration from the JSON file
	// if failed BotCfg is empty
	config.BotCfg = config.Read("../config.json")

	// WebSocket initialization
	websocket := Connection()

	// start listening on all channels
	Listen(websocket)

}

// print details of an error
func PrintError(err *model.AppError) {
	println("\tError Details:")
	println("\t\t" + err.Message)
	println("\t\t" + err.Id)
	println("\t\t" + err.DetailedError)
}

// connect with the Mattermost server
func Connection () *model.WebSocketClient{

	// create new MatterMost client
	config.MmCfg.Client = model.NewAPIv4Client("http://"+config.BotCfg.Server)

	// test to see if the mattermost server is up and running
	MakeSureServerIsRunning()

	// attempt to login to the Mattermost server as the bot user
	// This will set the token required for all future calls
	// You can get this token with client.AuthToken
	LoginAsTheBotUser()

	// find the bot team
	FindBotTeam()

	// create new WebSocket client
	var err *model.AppError
	config.MmCfg.WebSocketClient, err = model.NewWebSocketClient4("ws://"+config.BotCfg.Server, config.MmCfg.Client.AuthToken)
	if err != nil {
		println("We failed to connect to the web socket")
		PrintError(err)
	}

	return config.MmCfg.WebSocketClient
}

// check the mattermost server
func MakeSureServerIsRunning() {
	if props, resp := config.MmCfg.Client.GetOldClientConfig(""); resp.Error != nil {
		println("There was a problem pinging the Mattermost server.  Are you sure it's running?")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		// log
		println("Server detected and is running version " + props["Version"])
	}
}

// login to the chat as bot user using credentials from config
func LoginAsTheBotUser() {
	if user, resp := config.MmCfg.Client.Login(config.BotCfg.Email, config.BotCfg.Password); resp.Error != nil {
		println("There was a problem logging into the Mattermost server.  Are you sure ran the setup steps from the README.md?")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		config.MmCfg.BotUser = user
	}
}

// check whether the bot user is a member of the team specified in config
func FindBotTeam() {
	if team, resp := config.MmCfg.Client.GetTeamByName(config.BotCfg.TeamName, ""); resp.Error != nil {
		println("We failed to get the initial load")
		println("or we do not appear to be a member of the team '" + config.BotCfg.TeamName + "'")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		config.MmCfg.BotTeam = team
	}
}

// listen on all channels and start a bot (botVM)
func Listen(ws *model.WebSocketClient){
	ws.Listen()
	bot.Start(ws)
}