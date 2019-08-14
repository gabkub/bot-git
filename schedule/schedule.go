package schedule

import (
	"bot-git/bot/limit"
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
	_, e = scheduler.Every().Day().At("00:00").Run(resetPages)

	//_,e = scheduler.Every(config.DbCfg.ConnectionsCheckCron).Hours().Run(pgMonitor.CheckConnections)
	//_,e = scheduler.Every(config.DbCfg.ConnectionsLogCron).Hours().Run(pgMonitor.LogConnections)
	if e != nil {
		logg.WriteToFile("Error while scheduling. Details: " + e.Error())
	}
}

func resetRequests() {
	limit.Reset()
}

var pagesList = []map[string]int{
	newsSrc.GamePage, newsSrc.MediaPage, newsSrc.MotoPage, newsSrc.SciencePage, newsSrc.TechPage, newsSrc.VoyagePage,
}

func resetPages() {
	for _, list := range pagesList {
		for k, _ := range list {
			list[k] = 0
		}
	}
}
