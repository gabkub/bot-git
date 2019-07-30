package newsSrc

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc/newsAbstract"
)

var GetScience = []newsAbstract.GetNews{
	scienceSpider,
}

var sciencePage = map[string]int{
	"Spider": 0,
}
func scienceSpider() []messages.Message{
	sciencePage["Spider"]++
	return newsAbstract.GetSpider("nauka", sciencePage["Spider"])
}
