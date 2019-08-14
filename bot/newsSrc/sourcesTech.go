package newsSrc

import (
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var GetTech = []newsAbstract.GetNews{
	techSpider,
	techComputerWorld,
	techWirtualneMedia,
}

var TechPage = map[string]int{
	"Spider":         0,
	"ComputerWorld":  0,
	"WirtualneMedia": 0,
}

func techSpider() (*newsAbstract.News, bool) {
	return getSpider("nowe-technologie")
}

func techWirtualneMedia() (*newsAbstract.News, bool) {
	return getWirtualneMedia("technologie")
}

func techComputerWorld() (*newsAbstract.News, bool) {
	fetch := func(page int) *goquery.Selection {
		return contentFetcher.Fetch(fmt.Sprintf("https://www.computerworld.pl/news/archiwum-%v.html", TechPage["ComputerWorld"]), "div.row-list-item")
	}
	link := func(s *goquery.Selection) string {
		textLink, _ := s.Find("div.row-item-icon > a").Attr("href")
		return fmt.Sprintf("https://www.computerworld.pl/%v", textLink)
	}
	imgSel := func(s *goquery.Selection) string {
		image, _ := s.Find("div.row-item-icon > a > figure.frame-responsive img.img-responsive").Attr("src")
		return image
	}

	imgTextSel := func(s *goquery.Selection) string {
		return s.Find("div.col-lg-9 > a > span.title").Text()
	}

	sel := newSelectors(link, imgSel, imgTextSel)
	return getFreshNews(fetch, 5, sel)
}
