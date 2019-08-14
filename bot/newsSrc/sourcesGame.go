package newsSrc

import (
	"bot-git/bot/blacklist"
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var newsBlacklist = blacklist.New(10)

func Clean() {
	newsBlacklist.Clean()
}

var GetGame = []newsAbstract.GetNews{
	gameSpider, gamePPE,
}

func gameSpider() (*newsAbstract.News, bool) {
	return getSpider("gry")
}

func gamePPE() (*newsAbstract.News, bool) {

	linkBuilder := func(s *goquery.Selection) string {
		titleLink, _ := s.Find("div.txt > div.image_big > a.imgholder").Attr("href")
		return fmt.Sprintf("https://www.ppe.pl%v", titleLink)
	}

	fetch := func(page int) *goquery.Selection {
		return contentFetcher.Fetch(fmt.Sprintf("https://www.ppe.pl/news/news.html?page=%v", page), "div.box")
	}

	imgSel := func(s *goquery.Selection) string {
		i, _ := s.Find("div.txt div.image_big > a.imgholder > img.imgholderimg").Attr("src")
		return i
	}
	imgTextSel := func(s *goquery.Selection) string {
		t, _ := s.Find("div.txt div.image_big > a.imgholder > img.imgholderimg").Attr("alt")
		return t
	}
	sel := newSelectors(linkBuilder, imgSel, imgTextSel)
	return getFreshNews(fetch, 5, sel)
}
