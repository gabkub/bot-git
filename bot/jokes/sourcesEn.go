package jokes

import (
	"bot-git/bot/blacklists"
	"bot-git/contentFetcher"
	"fmt"
)

//var jokersEn = []getJoke{}
var jokersEn = []getJoke{
	iCanHazDadJoke,
	rd,
}

var countersEn = map[string]int{
	"rd": 1,
}

func iCanHazDadJoke() []string {
	blacklists.New("DadJokeBL")
	var jokes []string
	for i := 0; i < 10; i++ {
		div := contentFetcher.Fetch("https://icanhazdadjoke.com/", "div.card-content p")
		jokes = append(jokes, getJokesList(div)[0])
	}
	return jokes
}

func rd() []string {
	blacklists.New("rdBL")
	div := contentFetcher.Fetch(fmt.Sprintf("https://www.rd.com/jokes/page/%v/", countersEn["rd"]), "div.excerpt-wrapper")
	countersEn["rd"]++
	return getJokesList(div)
}
