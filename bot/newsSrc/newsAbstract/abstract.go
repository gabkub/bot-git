package newsAbstract

import (
	"bot-git/bot/messages"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type GetNews func() []messages.Message

func GetSpider(category string, page int) []messages.Message {
	//blacklists.New(fmt.Sprintf("%v",category))
	var messagesToReturn []messages.Message
	div := contentFetcher.Fetch(fmt.Sprintf("https://www.spidersweb.pl/kategoria/%v/page/%v", category, page), "div.columns:first-child div.columns")
	div.Each(func(i int, s *goquery.Selection) {
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
	var news []messages.Message

	div := contentFetcher.Fetch(fmt.Sprintf("https://www.wirtualnemedia.pl/wiadomosci/%v/page:%v", category, page), "div.news-box-content")
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
