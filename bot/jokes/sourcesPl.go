package jokes

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

var jokerPl = []getJoke{
	Perelki,
}

func Perelki() string {
	doc, err := goquery.NewDocument("https://perelki.net/random")

	if err != nil{
		return ""
	}

	div := doc.Find("div.container:first-child")

	if div == nil{
		return ""
	}

	result := div.Text()
	result = strings.ReplaceAll(div.Text(), doc.Find("div.about").Text(), "")
	result = strings.TrimSpace(result)

	return result
}

/*func Gomeo() string {
	doc, err := goquery.NewDocument("http://humor.gomeo.pl/krotkie-dowcipy")

	if err != nil{
		return ""
	}


	div := doc.Find("div.row div.joke-content headline-mag")

	if div == nil{
		return ""
	}

	result := div.Text()
	result = strings.ReplaceAll(div.Text(), doc.Find("div.about").Text(), "")
	result = strings.TrimSpace(result)

	return result
}*/
