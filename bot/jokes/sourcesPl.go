package jokes

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
	"fmt"
)

var jokersPl = []getJoke{
	jeja,
	gomeo,
}

var countersPl = map[string]int{
	"gomeo": 1,
}

func jeja() []string {
	blacklists.New("jejaBL")
	doc := abstract.GetDoc("https://dowcipy.jeja.pl/losowe")
	div := abstract.GetDiv(doc, "div.dow-left-text p")
	return getJokesList(div)
}

func gomeo() []string {
	blacklists.New("gomeoBL")
	doc := abstract.GetDoc(fmt.Sprintf("http://humor.gomeo.pl/krotkie-dowcipy/strona/%v", countersPl["gomeo"]))
	div := abstract.GetDiv(doc, "div.joke-content")
	countersPl["gomeo"]++
	return getJokesList(div)
}
