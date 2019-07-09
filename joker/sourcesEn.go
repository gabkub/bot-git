package joker

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

var jokerEn = []getJoke{
	ICanHazDadJoke,
}

func ICanHazDadJoke() string {
	doc, err := goquery.NewDocument("https://icanhazdadjoke.com/")

	if err != nil{
		return ""
	}

	//joke := FindRandom(doc,"div.",)

	div := doc.Find("div.card-content")

	if div == nil{
		return ""
	}

	result := div.Text()
	result = strings.TrimSpace(result)

	return result
}