package jokes

import (
	"bot-git/bot/blacklists"
	"bot-git/contentFetcher"
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
	div := contentFetcher.Fetch("https://dowcipy.jeja.pl/losowe", "div.dow-left-text p")
	return getJokesList(div)
}

func gomeo() []string {
	blacklists.New("gomeoBL")
	div := contentFetcher.Fetch(fmt.Sprintf("http://humor.gomeo.pl/krotkie-dowcipy/strona/%v", countersPl["gomeo"]), "div.joke-content")
	countersPl["gomeo"]++
	return getJokesList(div)
}
