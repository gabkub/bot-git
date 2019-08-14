package newsSrc

import (
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var GetVoyage = []newsAbstract.GetNews{
	voyageSpider,
	voyageMlecznePodroze,
}

var VoyagePage = map[string]int{
	"Spider":         0,
	"MlecznePodroze": 0,
}

func voyageSpider() (*newsAbstract.News, bool) {
	return getSpider("podroze")
}

func voyageMlecznePodroze() (*newsAbstract.News, bool) {
	fetch := func(page int) *goquery.Selection {
		return contentFetcher.Fetch(fmt.Sprintf("https://mlecznepodroze.pl/tag/news/page/%v/", VoyagePage["MlecznePodroze"]), "div.primary-post-content")
	}
	link := func(s *goquery.Selection) string {
		textLink, _ := s.Find("div.picture > div.picture-content > a").Attr("href")
		return textLink
	}
	imgSel := func(s *goquery.Selection) string {
		image, _ := s.Find("div.picture > div.picture-content > a > img").Attr("src")
		return image
	}

	imgTextSel := func(s *goquery.Selection) string {
		text, _ := s.Find("div.picture > div.picture-content > a").Attr("title")
		return text
	}
	sel := newSelectors(link, imgSel, imgTextSel)
	return getFreshNews(fetch, 5, sel)
}
