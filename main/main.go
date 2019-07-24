package main

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/logs"
	"github.com/mattermost/mattermost-bot-sample-golang/main/connection"
	"github.com/mattermost/mattermost-bot-sample-golang/schedule"
	"os"
	"os/signal"
)

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			if connection.Websocket != nil {
				connection.Websocket.Close()
			}
			logs.WriteToFile("Bot has exited the chat.")
			os.Exit(0)
		}
	}()
	os.Remove("./logs.log")
	logs.WriteToFile(fmt.Sprintf("Running bot v.%v...\n", commands.VER))
	connection.Connect()
	connection.Websocket.Listen()
	schedule.Start()
	bot.Start()
}

