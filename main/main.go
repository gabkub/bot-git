package main

import (
	"bot-git/bot"
	"bot-git/bot/commands/version"
	"bot-git/config"
	"bot-git/logg"
	"bot-git/main/connection"
	"fmt"
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
	logg.WriteToFile(fmt.Sprintf("Starting bot v.%v...\n", version.VER))

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
