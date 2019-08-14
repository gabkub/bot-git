package newsSrc

import (
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var GetScience = []newsAbstract.GetNews{
	sciencePrzystanek,
	scienceSpider,
}

var SciencePage = map[string]int{
	"Spider":     0,
	"Przystanek": -1,
}

func scienceSpider() (*newsAbstract.News, bool) {
	return getSpider("nauka")
}

func sciencePrzystanek() (*newsAbstract.News, bool) {
	fetch := func(page int) *goquery.Selection {
		return contentFetcher.Fetch(fmt.Sprintf("http://przystaneknauka.us.edu.pl/news?page=%v", SciencePage["Przystanek"]), "div.views-row")
	}
	link := func(s *goquery.Selection) string {
		textLink, _ := s.Find("h3.title > a").Attr("href")
		return fmt.Sprintf("http://przystaneknauka.us.edu.pl%v", textLink)
	}
	imgSel := func(s *goquery.Selection) string {
		image, _ := s.Find("a > img").Attr("src")
		return image
	}

	imgTextSel := func(s *goquery.Selection) string {
		return s.Find("h3.title").Text()
	}

	sel := newSelectors(link, imgSel, imgTextSel)
	return getFreshNews(fetch, 5, sel)
}
