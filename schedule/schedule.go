package schedule

import (
	"bot-git/bot/jokes"
	"bot-git/bot/limit"
	"bot-git/bot/memes"
	"bot-git/bot/newsSrc"
	"bot-git/logg"
	"github.com/carlescere/scheduler"
)

func Start() {
	_, e := scheduler.Every().Day().At("7:00").Run(resetRequests)
	_, e = scheduler.Every().Day().At("9:00").Run(resetRequests)
	_, e = scheduler.Every().Day().At("10:00").Run(resetRequests)
	_, e = scheduler.Every().Day().At("11:00").Run(resetRequests)
	_, e = scheduler.Every().Day().At("12:00").Run(resetRequests)
	_, e = scheduler.Every().Day().At("13:00").Run(resetRequests)
	_, e = scheduler.Every().Day().At("14:00").Run(resetRequests)
	_, e = scheduler.Every().Day().At("15:00").Run(resetRequests)
	_, e = scheduler.Every().Day().At("1:00").Run(cleanBlacklist)

	//_,e = scheduler.Every(config.DbCfg.ConnectionsCheckCron).Hours().Run(pgMonitor.CheckConnections)
	//_,e = scheduler.Every(config.DbCfg.ConnectionsLogCron).Hours().Run(pgMonitor.LogConnections)
	if e != nil {
		logg.WriteToFile("Error while scheduling. Details: " + e.Error())
	}
}

func cleanBlacklist() {
	memes.Blacklist.Clean()
	jokes.PolishBlacklist.Clean()
	jokes.EnglishJokesBlacklist.Clean()
	jokes.HardBlacklist.Clean()
	newsSrc.Clean()
}

func resetRequests() {
	limit.Reset()
}
