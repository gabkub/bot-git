package jokes

import (
	"bot-git/bot/blacklist"
	"bot-git/contentFetcher"
	"fmt"
)

var EnglishJokesBlacklist = blacklist.New(30, "english_jokes")
var jokersEn = []getJoke{
	iCanHazDadJoke,
	rd,
}

func iCanHazDadJoke() (*string, bool) {
	var jokes []*string
	for i := 0; i < 10; i++ {
		div := contentFetcher.Fetch("https://icanhazdadjoke.com/", "div.card-content p")
		jokes = append(jokes, getJokesList(div)[0])
	}
	return getFreshJoke(jokes, EnglishJokesBlacklist)
}

func rd() (*string, bool) {
	return getFreshForFetcher(fetchRd, 3, EnglishJokesBlacklist)
}

func fetchRd(page int) []*string {
	div := contentFetcher.Fetch(fmt.Sprintf("https://www.rd.com/jokes/page/%v/", page), "div.excerpt-wrapper")
	return getJokesList(div)
}
