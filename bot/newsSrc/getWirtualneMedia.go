package newsSrc

import (
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func getWirtualneMedia(category string) (*newsAbstract.News, bool) {
	fetch := func(page int) *goquery.Selection {
		return contentFetcher.Fetch(fmt.Sprintf("https://www.wirtualnemedia.pl/wiadomosci/%v/page:%v", category, page), "div.news-box-content")
	}
	link := func(s *goquery.Selection) string {
		textLink, _ := s.Find("div.news-img-wrapper > a").Attr("href")
		return fmt.Sprintf("https://www.wirtualnemedia.pl%v", textLink)
	}
	imgSel := func(s *goquery.Selection) string {
		image, _ := s.Find("div.news-img-wrapper > a > div.news-img-ratio > img").Attr("src")
		return image
	}

	imgTextSel := func(s *goquery.Selection) string {
		return s.Find("div.news-desc-head").Text()
	}

	sel := newSelectors(link, imgSel, imgTextSel)
	return getFreshNews(fetch, 5, sel)
}
