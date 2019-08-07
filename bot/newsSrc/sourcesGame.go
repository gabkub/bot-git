package newsSrc

import (
	"bot-git/bot/blacklists"
	"bot-git/bot/messages"
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var GetGame = []newsAbstract.GetNews{
	gameSpider, gamePPE,
}

var GamePage = map[string]int{
	"Spider": 0,
	"PPE":    0,
}

func gameSpider() []messages.Message {
	blacklists.New("gameSpiderBL")
	GamePage["Spider"]++
	return newsAbstract.GetSpider("gry", GamePage["Spider"])
}

func gamePPE() []messages.Message {
	blacklists.New("gamePPEBL")
	var messagesToReturn []messages.Message

	div := contentFetcher.Fetch(fmt.Sprintf("https://www.ppe.pl/news/news.html?page=%v", GamePage["PPE"]), "div.box")
	div.Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("div.txt div.image_big > a.imgholder > img.imgholderimg").Attr("src")
		text, _ := s.Find("div.txt div.image_big > a.imgholder > img.imgholderimg").Attr("alt")
		titleLink, _ := s.Find("div.txt > div.image_big > a.imgholder").Attr("href")
		message := messages.Message{
			TitleLink: fmt.Sprintf("https://www.ppe.pl%v", titleLink),
			Img: messages.Image{
				Header:   text,
				ImageUrl: image,
			},
		}
		if !message.Img.IsEmpty() && message.TitleLink != "" {
			messagesToReturn = append(messagesToReturn, message)
		}
	})
	return messagesToReturn
}
