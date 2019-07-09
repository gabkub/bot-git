package main

import (
	//"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-server/model"
	"os"
)

// Mattermost API connection data
var mmCfg config.MMConfig

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client
func main() {
	// This is an important step.  Lets make sure we use the botTeam
	// for all future web service requests that require a team.
	// client.SetTeamId(botTeam.Id)
	cfg, e := config.Read("config.json")

	if e != nil{
		//log
		println(e)
	}

	// inicjalizacja WebSocket
	socket := Connection(cfg)

	// rozpoczęcie nasłuchiwania na wszystkich kanałach
	Listen(socket, cfg)

}

// wypisanie szczegółów błędu
func PrintError(err *model.AppError) {
	println("\tError Details:")
	println("\t\t" + err.Message)
	println("\t\t" + err.Id)
	println("\t\t" + err.DetailedError)
}

// sprawdzenie
func MakeSureServerIsRunning() {
	if props, resp := mmCfg.Client.GetOldClientConfig(""); resp.Error != nil {
		println("There was a problem pinging the Mattermost server.  Are you sure it's running?")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		// log
		println("Server detected and is running version " + props["Version"])
	}
}

func LoginAsTheBotUser(cfg *config.BotConfig) {
	if user, resp := mmCfg.Client.Login(cfg.Email, cfg.Password); resp.Error != nil {
		println("There was a problem logging into the Mattermost server.  Are you sure ran the setup steps from the README.md?")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		mmCfg.BotUser = user
	}
}

func FindBotTeam(cfg *config.BotConfig) {
	if team, resp := mmCfg.Client.GetTeamByName(cfg.TeamName, ""); resp.Error != nil {
		println("We failed to get the initial load")
		println("or we do not appear to be a member of the team '" + cfg.TeamName + "'")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		mmCfg.BotTeam = team
	}
}

func Connection(cfg *config.BotConfig) *model.WebSocketClient{

	mmCfg.Client = model.NewAPIv4Client(cfg.Server)

	// test to see if the mattermost server is up and running
	MakeSureServerIsRunning()

	// attempt to login to the Mattermost server as the bot user
	// This will set the token required for all future calls
	// You can get this token with client.AuthToken
	LoginAsTheBotUser(cfg)

	// Lets find our bot team
	FindBotTeam(cfg)

	// Lets start listening to some channels via the websocket!
	var err *model.AppError
	mmCfg.WebSocketClient, err = model.NewWebSocketClient4("ws://192.168.3.182:8065", mmCfg.Client.AuthToken)
	if err != nil {
		println("We failed to connect to the web socket")
		PrintError(err)
	}

	return mmCfg.WebSocketClient
}

func Listen(ws *model.WebSocketClient, botConfig *config.BotConfig){
	ws.Listen()

	bot.Start(ws,botConfig, &mmCfg)
}