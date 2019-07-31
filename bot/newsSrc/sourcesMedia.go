package newsSrc

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc/newsAbstract"
)

var GetMedia = []newsAbstract.GetNews{
	mediaWirtualneMedia,
	mediaSpider,
}

var mediaPage = map[string]int{
	"Spider": 0,
	"WirtualneMedia":0,
}

func mediaSpider() []messages.Message{
	blacklists.New("mediaSpiderBL")
	mediaPage["Spider"]++
	return newsAbstract.GetSpider("media", mediaPage["Spider"])
}

func mediaWirtualneMedia() []messages.Message{
	blacklists.New("mediaWirtualneMediaBL")
	mediaPage["WirtualneMedia"]++
	return newsAbstract.GetWirtualneMedia("kultura-i-rozrywka", mediaPage["WirtualneMedia"])
}
