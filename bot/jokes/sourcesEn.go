package jokes

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklist"
)

var jokerEn = []getJoke{
	iCanHazDadJoke,
}

func iCanHazDadJoke() string {
	blacklist.New("DadJokeBL")
	doc := abstract.GetDoc("https://icanhazdadjoke.com/")
	div := abstract.GetDiv(doc, "div.card-content")

	result := fixFormat(div.Text())

	handleBlacklist(iCanHazDadJoke, result)

	return result
}