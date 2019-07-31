package main

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/logg"
	"github.com/mattermost/mattermost-bot-sample-golang/main/connection"
	"os"
	"os/signal"
	"syscall"
)

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client

func main() {
	setGracefulShutdown()
	logg.SetOutPut()
	config.ReadConfig()

	os.Remove("./logg.log")
	logg.WriteToFile(fmt.Sprintf("Starting bot v.%v...\n", commands.VER))

	connection.Connect()
	bot.Start()
}

func setGracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT)
	go func() {
		for range c {
			if connection.Websocket != nil {
				connection.Websocket.Close()
			}
			logg.WriteToFile("Bot has exited the chat.")
			os.Exit(0)
		}
	}()
}