package schedule

import (
	"github.com/carlescere/scheduler"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/logs"
)

func Start() {
	_,e := scheduler.Every().Day().At("7:00").Run(resetRequests)
	_,e = scheduler.Every().Day().At("9:00").Run(resetRequests)
	_,e = scheduler.Every().Day().At("10:00").Run(resetRequests)
	_,e = scheduler.Every().Day().At("11:00").Run(resetRequests)
	_,e = scheduler.Every().Day().At("12:00").Run(resetRequests)
	_,e = scheduler.Every().Day().At("13:00").Run(resetRequests)
	_,e = scheduler.Every().Day().At("14:00").Run(resetRequests)
	_,e = scheduler.Every().Day().At("15:00").Run(resetRequests)
	if e != nil {
		logs.WriteToFile("Error reseting user requests. Details: " + e.Error())
	}
}

func resetRequests() {
	for _,user := range limit.Users {
		for _,userLimit := range user {
			userLimit.Count = 0
			userLimit.LimitReached = false
		}
	}
}
