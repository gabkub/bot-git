package newsAbstract

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
)

type GetNews func() []messages.Message

func GetSpider(category string, page int) []messages.Message{
	//blacklists.New(fmt.Sprintf("%v",category))
	doc := abstract.GetDoc(fmt.Sprintf("https://www.spidersweb.pl/kategoria/%v/page/%v", category, page))
	var messagesToReturn []messages.Message
	abstract.GetDiv(doc,"div.columns:first-child div.columns").Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("article.article > div.cover > a.single-permalink > img").Attr("data-src")
		link, _ := s.Find("article.article > div.cover > a.single-permalink").Attr("href")
		message := messages.Message{
			TitleLink: link,
			Img: messages.Image{
				Header: s.Find("div.inner h1.title").Text(),
				ImageUrl: image,
			},
		}

		if !message.Img.IsEmpty() && message.TitleLink != ""{
			messagesToReturn = append(messagesToReturn, message)
		}
	})

	return messagesToReturn
}
