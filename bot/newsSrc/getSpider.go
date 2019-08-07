package newsSrc

import (
	"bot-git/bot/abstract"
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func getSpider(category string, page int) []*newsAbstract.News {
	//blacklists.New(fmt.Sprintf("%v",category))
	var messagesToReturn []*newsAbstract.News
	div := contentFetcher.Fetch(fmt.Sprintf("https://www.spidersweb.pl/kategoria/%v/page/%v", category, page), "div.columns:first-child div.columns")
	div.Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("article.article > div.cover > a.single-permalink > img").Attr("data-src")
		link, _ := s.Find("article.article > div.cover > a.single-permalink").Attr("href")
		img := abstract.NewImage(s.Find("div.inner h1.title").Text(), image)
		message := newsAbstract.NewNews(link, img)
		if !message.Img.IsEmpty() && message.TitleLink != "" {
			messagesToReturn = append(messagesToReturn, message)
		}
	})

	return messagesToReturn
}
