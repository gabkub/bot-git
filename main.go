package main

import (
	"./bot"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-server/model"
	"os"
)

// Mattermost API connection data
/*var client *model.Client4
var webSocketClient *model.WebSocketClient
var botUser *model.User
var botTeam *model.TeamName
var debuggingChannel *model.Channel*/
var mmCfg bot.MMConfig

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client
func main() {
	// This is an important step.  Lets make sure we use the botTeam
	// for all future web service requests that require a team.
	// client.SetTeamId(botTeam.Id)
	cfg, e := config.Read("config.json")
	println(cfg.Name, cfg.Password, cfg.Email, cfg.TeamName, cfg.EnglishDay)
	if e != nil{
		//log
		println(e)
	}
	socket := Connection(cfg)

	Listen(socket, cfg)

}

func PrintError(err *model.AppError) {
	println("\tError Details:")
	println("\t\t" + err.Message)
	println("\t\t" + err.Id)
	println("\t\t" + err.DetailedError)
}

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

	mmCfg.Client = model.NewAPIv4Client("http://192.168.3.182:8065")

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

	go func() {
		for {
			select {
			case resp := <-ws.EventChannel:
				bot.Start(resp, &mmCfg, botConfig)
			}
		}
	}()
	// You can block forever with
	select {}
}

/*func CreateBotDebuggingChannelIfNeeded() {
	if rchannel, resp := client.GetChannelByName(CHANNEL_LOG_NAME, botTeam.Id, ""); resp.Error != nil {
		println("We failed to get the channels")
		PrintError(resp.Error)
	} else {
		debuggingChannel = rchannel
		return
	}

	// Looks like we need to create the logging channel
	channel := &model.Channel{}
	channel.Name = CHANNEL_LOG_NAME
	channel.DisplayName = "Debugging For Sample Bot"
	channel.Purpose = "This is used as a test channel for logging bot debug messages"
	channel.Type = model.CHANNEL_OPEN
	channel.TeamId = botTeam.Id
	if rchannel, resp := client.CreateChannel(channel); resp.Error != nil {
		println("We failed to create the channel " + CHANNEL_LOG_NAME)
		PrintError(resp.Error)
	} else {
		debuggingChannel = rchannel
		println("Looks like this might be the first run so we've created the channel " + CHANNEL_LOG_NAME)
	}
}

func SetupGracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			if webSocketClient != nil {
				webSocketClient.Close()
			}

			//SendMsg("_"+SAMPLE_NAME+" has **stopped** running_", "")
			os.Exit(0)
		}
	}()
}*/
