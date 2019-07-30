package newsSrc

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc/newsAbstract"
)

var GetMedia = []newsAbstract.GetNews{
	mediaSpider,
}

var mediaPage = map[string]int{
	"Spider": 0,
}

func mediaSpider() []messages.Message{
	mediaPage["Spider"]++
	return newsAbstract.GetSpider("media", mediaPage["Spider"])
}
