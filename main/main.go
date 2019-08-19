package main

import (
	"bot-git/bot"
	"bot-git/bot/commands/version"
	"bot-git/config"
	"bot-git/main/connection"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client

func main() {
	setGracefulShutdown()
	config.ReadConfig()

	log.Println(fmt.Sprintf("Starting bot v.%v...\n", version.VER))

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
			log.Println("Bot has exited the chat.")
			os.Exit(0)
		}
	}()
}

func init() {
	fileName := fmt.Sprintf("./log_%s.log", time.Now().Format("2019-06-12"))
	createFile(fileName)
	log.SetOutput(&lumberjack.Logger{
		Filename: fileName,
		MaxSize:  1000,
		MaxAge:   7,
	})
}
func createFile(fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()
}
