package newsSrc

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc/newsAbstract"
)

var GetVoyage = []newsAbstract.GetNews{
	voyageSpider,
}

var voyagePage = map[string]int{
	"Spider": 0,
}
func voyageSpider() []messages.Message{
	voyagePage["Spider"]++
	return newsAbstract.GetSpider("podroze", voyagePage["Spider"])
}
