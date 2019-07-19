package main

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/main/connection"
	"log"
	"os"
	"os/signal"
)

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client

func main() {
	log.Printf("Running bot v.%v...\n", commands.VER)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			if connection.Websocket != nil {
				connection.Websocket.Close()
			}
			log.Println("Bot has exited the chat.")
			os.Exit(0)
		}
	}()
	// WebSocket initialization
	connection.Connection()

	// start listening on all channels

	bot.Start(connection.Websocket)
}

