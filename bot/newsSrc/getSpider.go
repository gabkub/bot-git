package newsSrc

import (
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func getSpider(category string) (*newsAbstract.News, bool) {
	fetch := func(page int) *goquery.Selection {
		return contentFetcher.Fetch(fmt.Sprintf("https://www.spidersweb.pl/kategoria/%v/page/%v", category, page), "div.columns:first-child div.columns")
	}
	link := func(s *goquery.Selection) string {
		l, _ := s.Find("article.article > div.cover > a.single-permalink").Attr("href")
		return l
	}
	imgSel := func(s *goquery.Selection) string {
		l, _ := s.Find("article.article > div.cover > a.single-permalink > img").Attr("data-src")
		return l
	}

	imgTextSel := func(s *goquery.Selection) string {
		return s.Find("div.inner h1.title").Text()
	}

	sel := newSelectors(link, imgSel, imgTextSel)
	return getFreshNews(fetch, 5, sel)
}
