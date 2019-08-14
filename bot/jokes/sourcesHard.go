package jokes

import (
	"bot-git/bot/blacklist"
	"bot-git/contentFetcher"
)

var HardBlacklist = blacklist.New(30)

var jokersHard = []getJoke{
	suchary,
}

func suchary() (*string, bool) {
	div := contentFetcher.Fetch("http://suchary.jakubchmura.pl/obcy/random/", "div.panel-body p")
	jokes := getJokesList(div)
	return getFreshJoke(jokes, HardBlacklist)
}
