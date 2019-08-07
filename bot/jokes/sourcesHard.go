package jokes

import (
	"bot-git/bot/blacklists"
	"bot-git/contentFetcher"
)

var jokersHard = []getJoke{
	suchary,
}

var countersHard = map[string]int{}

func suchary() []string {
	blacklists.New("sucharyBL")
	div := contentFetcher.Fetch("http://suchary.jakubchmura.pl/obcy/random/", "div.panel-body p")
	countersHard["suchary"]++
	return getJokesList(div)
}
