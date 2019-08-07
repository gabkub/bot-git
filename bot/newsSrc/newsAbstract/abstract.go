package newsAbstract

import (
	"bot-git/bot/abstract"
	"bot-git/bot/messages"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type GetNews func() []messages.Message

func GetSpider(category string, page int) []messages.Message {
	//blacklists.New(fmt.Sprintf("%v",category))
	doc := abstract.GetDoc(fmt.Sprintf("https://www.spidersweb.pl/kategoria/%v/page/%v", category, page))
	var messagesToReturn []messages.Message
	abstract.GetDiv(doc, "div.columns:first-child div.columns").Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("article.article > div.cover > a.single-permalink > img").Attr("data-src")
		link, _ := s.Find("article.article > div.cover > a.single-permalink").Attr("href")
		message := messages.Message{
			TitleLink: link,
			Img: messages.Image{
				Header:   s.Find("div.inner h1.title").Text(),
				ImageUrl: image,
			},
		}

		if !message.Img.IsEmpty() && message.TitleLink != "" {
			messagesToReturn = append(messagesToReturn, message)
		}
	})

	return messagesToReturn
}

func GetWirtualneMedia(category string, page int) []messages.Message {
	doc := abstract.GetDoc(fmt.Sprintf("https://www.wirtualnemedia.pl/wiadomosci/%v/page:%v", category, page))

	div := abstract.GetDiv(doc, "div.news-box-content")

	var news []messages.Message

	div.Each(func(i int, s *goquery.Selection) {

		image, _ := s.Find("div.news-img-wrapper > a > div.news-img-ratio > img").Attr("src")
		text := s.Find("div.news-desc-head").Text()
		textlink, _ := s.Find("div.news-img-wrapper > a").Attr("href")
		temp := messages.Message{
			TitleLink: fmt.Sprintf("https://www.wirtualnemedia.pl%v", textlink),
			Img: messages.Image{
				Header:   text,
				ImageUrl: image,
			},
		}

		if !temp.Img.IsEmpty() && temp.TitleLink != "" {
			news = append(news, temp)
		}
	})
	return news
}
