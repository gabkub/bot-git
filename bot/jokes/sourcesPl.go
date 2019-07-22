package jokes

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
	"strings"
)

var jokerPl = []getJoke{
	perelki,
}

func perelki() string {
	blacklists.New("perelkiBL")

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
}/*
func jeja() []string{
	resp, err := http.Get("https://dowcipy.jeja.pl/losowe")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	var jokes []string
	doc.Find("div.dow-left-text:first-child").Each(func(i int, s *goquery.Selection) {
		println(strings.TrimSpace(s.Text()))
		jokes = append(jokes, s.Text())
	})

	return jokes
}
*/