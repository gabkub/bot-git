package main

import (
	"./bot"
	"github.com/mattermost/mattermost-server/model"
	"os"
)

// authorization data of bot's user
const (
	USER_EMAIL    = "bot2@example.com"
	USER_PASSWORD = "password1"
	/*USER_NAME     = "samplebot2"
	USER_FIRST    = "Sample"
	USER_LAST     = "Bot"*/
	TEAM_NAME        = "Test"
	//CHANNEL_LOG_NAME = "Off-Topic"
)

// Mattermost API connection data
/*var client *model.Client4
var webSocketClient *model.WebSocketClient
var botUser *model.User
var botTeam *model.Team
var debuggingChannel *model.Channel*/
var botCfg bot.BotConfig

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client
func main() {
	// This is an important step.  Lets make sure we use the botTeam
	// for all future web service requests that require a team.
	// client.SetTeamId(botTeam.Id)

	Connection()

}

func PrintError(err *model.AppError) {
	println("\tError Details:")
	println("\t\t" + err.Message)
	println("\t\t" + err.Id)
	println("\t\t" + err.DetailedError)
}

func MakeSureServerIsRunning() {
	if props, resp := botCfg.Client.GetOldClientConfig(""); resp.Error != nil {
		println("There was a problem pinging the Mattermost server.  Are you sure it's running?")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		// log
		println("Server detected and is running version " + props["Version"])
	}
}

func LoginAsTheBotUser() {
	if user, resp := botCfg.Client.Login(USER_EMAIL, USER_PASSWORD); resp.Error != nil {
		println("There was a problem logging into the Mattermost server.  Are you sure ran the setup steps from the README.md?")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		botCfg.BotUser = user
	}
}

func FindBotTeam() {
	if team, resp := botCfg.Client.GetTeamByName(TEAM_NAME, ""); resp.Error != nil {
		println("We failed to get the initial load")
		println("or we do not appear to be a member of the team '" + TEAM_NAME + "'")
		PrintError(resp.Error)
		os.Exit(1)
	} else {
		botCfg.BotTeam = team
	}
}

func Connection(){

	botCfg.Client = model.NewAPIv4Client("http://192.168.3.182:8065")

	// test to see if the mattermost server is up and running
	MakeSureServerIsRunning()

	// attempt to login to the Mattermost server as the bot user
	// This will set the token required for all future calls
	// You can get this token with client.AuthToken
	LoginAsTheBotUser()

	// Lets find our bot team
	FindBotTeam()

	// Lets start listening to some channels via the websocket!
	var err *model.AppError
	botCfg.WebSocketClient, err = model.NewWebSocketClient4("ws://192.168.3.182:8065", botCfg.Client.AuthToken)
	if err != nil {
		println("We failed to connect to the web socket")
		PrintError(err)
	}


	Listen(botCfg.WebSocketClient)
}

func Listen(ws *model.WebSocketClient){
	ws.Listen()

	go func() {
		for {
			select {
			case resp := <-ws.EventChannel:
				bot.Start(resp, &botCfg)
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
