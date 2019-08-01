package jokes

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
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
