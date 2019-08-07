package newsSrc

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var GetMoto = []newsAbstract.GetNews{
	motoAutoCentrum,
}

var MotoPage = map[string]int{
	"AutoCentrum": 0,
	"Moto":        0,
}

func motoAutoCentrum() []*newsAbstract.News {
	blacklists.New("motoAutoCentrumBL")
	MotoPage["AutoCentrum"]++
	var news []*newsAbstract.News

	div := contentFetcher.Fetch(fmt.Sprintf("https://www.autocentrum.pl/newsy/strona-%v/", MotoPage["AutoCentrum"]), "div.ac-article-wrapper")
	div.Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("a > div.photo > picture > img.img-responsive").Attr("src")
		text, _ := s.Find("a > div.photo > picture > img.img-responsive").Attr("alt")
		textLink, _ := s.Find("a.news-box").Attr("href")
		link := fmt.Sprintf("https://www.autocentrum.pl%v", textLink)
		img := abstract.NewImage(text, fmt.Sprintf("https://www.autocentrum.pl%v", image))
		temp := newsAbstract.NewNews(link, img)
		if !temp.Img.IsEmpty() && temp.TitleLink != "" {
			news = append(news, temp)
		}
	})
	return news
}

// Todo source not used probable scraping needs tuning
//func motoMoto() []messages.Message {
//	MotoPage["Moto"]++
//	var news []messages.Message
//	div := contentFetcher.Fetch(fmt.Sprintf("http://moto.pl/MotoPL/0,88389.html?str=%v_24561775", MotoPage["Moto"]), "li.entry")
//	div.Each(func(i int, s *goquery.Selection) {
//		image, _ := s.Find("a > fiugre.imgw > img").Attr("src")
//		text, _ := s.Find("a").Attr("title")
//		textlink, _ := s.Find("a").Attr("href")
//		temp := messages.Message{
//			TitleLink: textlink,
//			Img: messages.Image{
//				Header:   text,
//				ImageUrl: image,
//			},
//		}
//		if !temp.Img.IsEmpty() && temp.TitleLink != "" {
//			news = append(news, temp)
//		}
//	})
//	return news
//}
