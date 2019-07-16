package limit

import (
	"github.com/robfig/cron"
)

func Start() {
	c := cron.New()
	c.AddFunc("0 7,9-15 * * 1-5", resetRequests)
	c.Start()
}

func resetRequests() {
	for _,user := range Users {
		for _,limit := range user {
			limit.Count = 0
			limit.LimitReached = false
		}
	}
}
