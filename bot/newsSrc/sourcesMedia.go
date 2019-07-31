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

var MediaPage = map[string]int{
	"Spider": 0,
	"WirtualneMedia":0,
}

func mediaSpider() []messages.Message{
	blacklists.New("mediaSpiderBL")
	MediaPage["Spider"]++
	return newsAbstract.GetSpider("media", MediaPage["Spider"])
}

func mediaWirtualneMedia() []messages.Message{
	blacklists.New("mediaWirtualneMediaBL")
	MediaPage["WirtualneMedia"]++
	return newsAbstract.GetWirtualneMedia("kultura-i-rozrywka", MediaPage["WirtualneMedia"])
}
