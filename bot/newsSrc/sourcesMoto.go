package newsSrc

import (
	"bot-git/bot/blacklists"
	"bot-git/bot/messages"
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

func motoAutoCentrum() []messages.Message {
	blacklists.New("motoAutoCentrumBL")
	MotoPage["AutoCentrum"]++
	var news []messages.Message

	div := contentFetcher.Fetch(fmt.Sprintf("https://www.autocentrum.pl/newsy/strona-%v/", MotoPage["AutoCentrum"]), "div.ac-article-wrapper")
	div.Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("a > div.photo > picture > img.img-responsive").Attr("src")
		text, _ := s.Find("a > div.photo > picture > img.img-responsive").Attr("alt")
		textlink, _ := s.Find("a.news-box").Attr("href")
		temp := messages.Message{
			TitleLink: fmt.Sprintf("https://www.autocentrum.pl%v", textlink),
			Img: messages.Image{
				Header:   text,
				ImageUrl: fmt.Sprintf("https://www.autocentrum.pl%v", image),
			},
		}
		if !temp.Img.IsEmpty() && temp.TitleLink != "" {
			news = append(news, temp)
		}
	})
	return news
}
func motoMoto() []messages.Message {
	MotoPage["Moto"]++
	var news []messages.Message
	div := contentFetcher.Fetch(fmt.Sprintf("http://moto.pl/MotoPL/0,88389.html?str=%v_24561775", MotoPage["Moto"]), "li.entry")
	div.Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("a > fiugre.imgw > img").Attr("src")
		text, _ := s.Find("a").Attr("title")
		textlink, _ := s.Find("a").Attr("href")
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
