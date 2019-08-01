package newsSrc

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
	"bot-git/bot/messages"
	"bot-git/bot/newsSrc/newsAbstract"
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

func scienceSpider() []messages.Message {
	blacklists.New("scienceSpiderBL")
	SciencePage["Spider"]++
	return newsAbstract.GetSpider("nauka", SciencePage["Spider"])
}

func sciencePrzystanek() []messages.Message {
	blacklists.New("sciencePrzystanekBL")
	SciencePage["Przystanek"]++
	doc := abstract.GetDoc(fmt.Sprintf("http://przystaneknauka.us.edu.pl/news?page=%v", SciencePage["Przystanek"]))

	div := abstract.GetDiv(doc, "div.views-row")

	var news []messages.Message

	div.Each(func(i int, s *goquery.Selection) {

		image, _ := s.Find("a > img").Attr("src")
		text := s.Find("h3.title").Text()
		textlink, _ := s.Find("h3.title > a").Attr("href")
		temp := messages.Message{
			TitleLink: fmt.Sprintf("http://przystaneknauka.us.edu.pl%v", textlink),
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
