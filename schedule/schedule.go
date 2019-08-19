package schedule

import (
	"bot-git/bot/jokes"
	"bot-git/bot/limit"
	"bot-git/bot/memes"
	"bot-git/bot/newsSrc"
	"github.com/carlescere/scheduler"
)

func Start() {
	scheduler.Every().Day().At("7:00").Run(resetRequests)
	scheduler.Every().Day().At("9:00").Run(resetRequests)
	scheduler.Every().Day().At("10:00").Run(resetRequests)
	scheduler.Every().Day().At("11:00").Run(resetRequests)
	scheduler.Every().Day().At("12:00").Run(resetRequests)
	scheduler.Every().Day().At("13:00").Run(resetRequests)
	scheduler.Every().Day().At("14:00").Run(resetRequests)
	scheduler.Every().Day().At("15:00").Run(resetRequests)
	scheduler.Every().Day().At("1:00").Run(cleanBlacklist)
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
