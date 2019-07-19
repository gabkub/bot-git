package schedule

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"github.com/mattermost/mattermost-bot-sample-golang/bot"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/main/connection"
	"log"
)

func Start() {
	scheduler.Every().Day().At("7").NotImmediately().Run(resetRequests)
	scheduler.Every().Day().At("9").NotImmediately().Run(resetRequests)
	scheduler.Every().Day().At("10").NotImmediately().Run(resetRequests)
	scheduler.Every().Day().At("11").NotImmediately().Run(resetRequests)
	scheduler.Every().Day().At("12").NotImmediately().Run(resetRequests)
	scheduler.Every().Day().At("13").NotImmediately().Run(resetRequests)
	scheduler.Every().Day().At("14").NotImmediately().Run(resetRequests)
	scheduler.Every().Day().At("15").NotImmediately().Run(resetRequests)
	scheduler.Every(1).Minutes().NotImmediately().Run(checkConnection)
}

func resetRequests() {
	for _,user := range limit.Users {
		for _,userLimit := range user {
			userLimit.Count = 0
			userLimit.LimitReached = false
		}
	}
}

func checkConnection() {
	if ping, resp := config.MmCfg.Client.GetPing(); resp.Error != nil {
		log.Println("Server not responding. Connecting again.")
		connection.ConnectServer()
		bot.Start(connection.Websocket)
	} else {
		log.Println(fmt.Sprintf("Server ping: %v", ping))
	}
}

