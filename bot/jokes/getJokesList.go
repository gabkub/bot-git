package jokes

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func getJokesList(selectionsToFormat *goquery.Selection) []*string {
	var jokes []*string
	selectionsToFormat.Each(func(i int, s *goquery.Selection) {
		selectionHTML, _ := s.Html()
		jokes = append(jokes, fixFormat(selectionHTML))
	})
	return jokes
}

func fixFormat(HTMLtoFormat string) *string {
	formattedString := strings.ReplaceAll(HTMLtoFormat, "<br>", "\n")
	formattedString = strings.ReplaceAll(formattedString, "<br/>", "\n")
	formattedString = strings.ReplaceAll(formattedString, "<p>", "")
	formattedString = strings.ReplaceAll(formattedString, "</p>", "")

	// markdown escape
	formattedString = strings.ReplaceAll(formattedString, "-", "\\-")
	formattedString = strings.ReplaceAll(formattedString, "*", "\\*")

	formattedString = strings.TrimSpace(formattedString)

	return &formattedString
}
