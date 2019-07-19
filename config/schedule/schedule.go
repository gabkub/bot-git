package schedule

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
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
	scheduler.Every(10).Minutes().NotImmediately().Run(checkConnection)
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
		log.Fatal("Server not responding.")
	} else {
		log.Println(fmt.Sprintf("Server ping: %v", ping))
	}
}
