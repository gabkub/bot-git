package newsSrc

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc/newsAbstract"
)

var GetGame = []newsAbstract.GetNews{
	gameSpider,gamePPE,
}

var gamePage = map[string]int{
	"Spider": 0,
	"PPE": 0,
}

func gameSpider() []messages.Message{
	gamePage["Spider"]++
	return newsAbstract.GetSpider("gry", gamePage["Spider"])
}

func gamePPE() []messages.Message{
	doc := abstract.GetDoc(fmt.Sprintf("https://www.ppe.pl/news/news.html?page=%v",gamePage["PPE"]))
	var messagesToReturn []messages.Message
	abstract.GetDiv(doc,"div.box").Each(func(i int, s *goquery.Selection){

		image,_ := s.Find("div.txt div.image_big a.imgholder img.imgholderimg").Attr("src")
		text,_ := s.Find("div.txt div.image_big a.imgholder img.imgholderimg").Attr("alt")
		titleLink, _ := s.Find("div.txt div.image_big a.imgholder").Attr("href")
		 message := messages.Message{
			TitleLink:  titleLink,
			Img: messages.Image{
				Header: text,
				ImageUrl: image,
			},
		}

		 if !message.Img.IsEmpty() && message.TitleLink != ""{
		 	messagesToReturn = append(messagesToReturn, message)
		 }
	})
	return messagesToReturn
}
