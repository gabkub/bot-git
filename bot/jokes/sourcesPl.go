package jokes

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
)

var jokersPl = []getJoke{
	jeja,
	gomeo,
}

var countersPl = map[string]int {
	"gomeo": 1,
}

/* probably too extreme

func suchary() []string {
	blacklists.New("sucharyBL")
	doc := abstract.GetDoc(fmt.Sprintf("https://suchary.jakubchmura.pl/?page=%v", countersPl["suchary"]))
	div := abstract.GetDiv(doc, "div.panel-body p")
	countersPl["suchary"]++
	return getJokesList(div)
}*/

func jeja() []string {
	blacklists.New("jejaBL")
	doc := abstract.GetDoc("https://dowcipy.jeja.pl/losowe")
	div := abstract.GetDiv(doc, "div.dow-left-text p")
	return getJokesList(div)
}

func gomeo() []string {
	blacklists.New("gomeoBL")
	doc := abstract.GetDoc(fmt.Sprintf("http://humor.gomeo.pl/krotkie-dowcipy/strona/%v", countersPl["gomeo"]))
	div := abstract.GetDiv(doc,"div.joke-content")
	countersPl["gomeo"]++
	return getJokesList(div)
}