package schedule

import (
	"bot-git/bot/jokes"
	"bot-git/bot/limit"
	"bot-git/bot/memes"
	"bot-git/bot/newsSrc"
	"github.com/carlescere/scheduler"
	"log"
)

var schedulerRestRequestRuns = []string{"7:00", "9:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00"}

func Start() {
	for _, r := range schedulerRestRequestRuns {
		_, err := scheduler.Every().Day().At(r).Run(resetRequests)
		if err != nil {
			log.Println(err)
		}
	}
	_, err := scheduler.Every().Day().At("1:00").Run(cleanBlacklist)
	if err != nil {
		log.Println(err)
	}
}

func cleanBlacklist() {
	memes.Blacklist.Clean()
	jokes.PolishJokesBlacklist.Clean()
	jokes.EnglishJokesBlacklist.Clean()
	newsSrc.Clean()
}

func resetRequests() {
	limit.Reset()
}
