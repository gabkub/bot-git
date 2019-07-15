package jokes

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklist"
	"strings"
)

var jokerPl = []getJoke{
	perelki,
}

func perelki() string {
	blacklist.New("perelkiBL")

	doc := abstract.GetDoc("https://perelki.net/random")
	div := abstract.GetDiv(doc, "div.container:first-child")

	result := div.Text()
	result = strings.ReplaceAll(div.Text(), doc.Find("div.about").Text(), "")
	result = strings.TrimSpace(result)

	handleBL(perelki, result)

	return result
}
