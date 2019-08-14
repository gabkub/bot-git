package newsSrc

import (
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var GetMoto = []newsAbstract.GetNews{
	motoAutoCentrum,
}

func motoAutoCentrum() (*newsAbstract.News, bool) {
	fetch := func(page int) *goquery.Selection {
		return contentFetcher.Fetch(fmt.Sprintf("https://www.autocentrum.pl/newsy/strona-%v/", page), "div.ac-article-wrapper")
	}
	link := func(s *goquery.Selection) string {
		textLink, _ := s.Find("a.news-box").Attr("href")
		return fmt.Sprintf("https://www.autocentrum.pl%v", textLink)
	}
	imgSel := func(s *goquery.Selection) string {
		image, _ := s.Find("a > div.photo > picture > img.img-responsive").Attr("src")
		return fmt.Sprintf("https://www.autocentrum.pl%v", image)
	}

	imgTextSel := func(s *goquery.Selection) string {
		text, _ := s.Find("a > div.photo > picture > img.img-responsive").Attr("alt")
		return text
	}

	sel := newSelectors(link, imgSel, imgTextSel)
	return getFreshNews(fetch, 5, sel)
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
