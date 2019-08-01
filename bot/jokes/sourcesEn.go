package jokes

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
	"fmt"
)

//var jokersEn = []getJoke{}
var jokersEn = []getJoke{
	iCanHazDadJoke,
	rd,
}

var countersEn = map[string]int {
	"rd": 1,
}

func iCanHazDadJoke() []string {
	blacklists.New("DadJokeBL")
	var jokes []string
	for i:=0; i<10; i++ {
		doc := abstract.GetDoc("https://icanhazdadjoke.com/")
		div := abstract.GetDiv(doc, "div.card-content p")
		jokes = append(jokes, getJokesList(div)[0])
	}
	return jokes
}

func rd() []string {
	blacklists.New("rdBL")
	doc := abstract.GetDoc(fmt.Sprintf("https://www.rd.com/jokes/page/%v/", countersEn["rd"]))
	div := abstract.GetDiv(doc, "div.excerpt-wrapper")
	countersEn["rd"]++
	return getJokesList(div)
}

