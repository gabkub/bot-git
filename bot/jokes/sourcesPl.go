package jokes

import (
	"bot-git/bot/blacklist"
	"bot-git/contentFetcher"
	"fmt"
)

var PolishBlacklist = blacklist.New(30)
var jokersPl = []getJoke{
	jeja,
	gomeo,
}

func jeja() (*string, bool) {
	div := contentFetcher.Fetch("https://dowcipy.jeja.pl/losowe", "div.dow-left-text p")
	return getFreshJoke(getJokesList(div), PolishBlacklist)
}

func gomeo() (*string, bool) {
	return getFreshForFetcher(gomeoFetch, 4, PolishBlacklist)
}

func gomeoFetch(page int) []*string {
	div := contentFetcher.Fetch(fmt.Sprintf("http://humor.gomeo.pl/krotkie-dowcipy/strona/%v", page), "div.joke-content")
	return getJokesList(div)
}
