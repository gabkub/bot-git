package newsSrc

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
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

func gameSpider() []*newsAbstract.News {
	blacklists.New("gameSpiderBL")
	GamePage["Spider"]++
	return getSpider("gry", GamePage["Spider"])
}

func gamePPE() []*newsAbstract.News {
	blacklists.New("gamePPEBL")
	var messagesToReturn []*newsAbstract.News

	div := contentFetcher.Fetch(fmt.Sprintf("https://www.ppe.pl/news/news.html?page=%v", GamePage["PPE"]), "div.box")
	div.Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("div.txt div.image_big > a.imgholder > img.imgholderimg").Attr("src")
		text, _ := s.Find("div.txt div.image_big > a.imgholder > img.imgholderimg").Attr("alt")
		titleLink, _ := s.Find("div.txt > div.image_big > a.imgholder").Attr("href")
		link := fmt.Sprintf("https://www.ppe.pl%v", titleLink)
		img := abstract.NewImage(text, image)
		message := newsAbstract.NewNews(link, img)
		if !message.Img.IsEmpty() && message.TitleLink != "" {
			messagesToReturn = append(messagesToReturn, message)
		}
	})
	return messagesToReturn
}
