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

	div := abstract.GetDiv(doc, "div.content div.container:first-child")
	resultHTML, _ := div.Html()

	// cleaning the text
	toRemove, _ := doc.Find("div.about").Html()
	result := strings.ReplaceAll(resultHTML, toRemove, "")
	result = strings.ReplaceAll(result, "<div class=\"about\"></div>", "")
	result = fixFormat(result)

	handleBlacklist(perelki, result)

	return result
}