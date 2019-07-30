package jokes

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
)

var jokersHard = []getJoke{
	suchary,
}

var countersHard = map[string]int {
}

func suchary() []string {
	blacklists.New("sucharyBL")
	doc := abstract.GetDoc("http://suchary.jakubchmura.pl/obcy/random/")
	div := abstract.GetDiv(doc, "div.panel-body p")
	countersHard["suchary"]++
	return getJokesList(div)
}
