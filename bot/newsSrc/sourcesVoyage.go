package newsSrc

import (
	"bot-git/bot/blacklists"
	"bot-git/bot/messages"
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

func voyageSpider() []messages.Message {
	blacklists.New("voyageSpiderBL")
	VoyagePage["Spider"]++
	return newsAbstract.GetSpider("podroze", VoyagePage["Spider"])
}

func voyageMlecznePodroze() []messages.Message {
	blacklists.New("voyageMlecznePodrozeBL")
	VoyagePage["MlecznePodroze"]++
	var news []messages.Message
	div := contentFetcher.Fetch(fmt.Sprintf("https://mlecznepodroze.pl/tag/news/page/%v/", VoyagePage["MlecznePodroze"]), "div.primary-post-content")
	div.Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("div.picture > div.picture-content > a > img").Attr("src")
		text, _ := s.Find("div.picture > div.picture-content > a").Attr("title")
		textlink, _ := s.Find("div.picture > div.picture-content > a").Attr("href")
		temp := messages.Message{
			TitleLink: textlink,
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
