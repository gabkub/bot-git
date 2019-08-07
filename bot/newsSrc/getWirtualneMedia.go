package newsSrc

import (
	"bot-git/bot/abstract"
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func getWirtualneMedia(category string, page int) []*newsAbstract.News {
	var news []*newsAbstract.News

	div := contentFetcher.Fetch(fmt.Sprintf("https://www.wirtualnemedia.pl/wiadomosci/%v/page:%v", category, page), "div.news-box-content")
	div.Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("div.news-img-wrapper > a > div.news-img-ratio > img").Attr("src")
		text := s.Find("div.news-desc-head").Text()
		textLink, _ := s.Find("div.news-img-wrapper > a").Attr("href")
		title := fmt.Sprintf("https://www.wirtualnemedia.pl%v", textLink)
		img := abstract.NewImage(text, image)
		temp := newsAbstract.NewNews(title, img)
		if !temp.Img.IsEmpty() && temp.TitleLink != "" {
			news = append(news, temp)
		}
	})
	return news
}
