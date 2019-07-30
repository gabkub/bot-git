package newsSrc

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc/newsAbstract"
)

var GetTech = []newsAbstract.GetNews{
	techComputerWorld,
}

var techPage = map[string]int{
	"Spider": 0,
	"ComputerWorld": 0,
}

func techSpider() []messages.Message{
	techPage["Spider"]++
	return newsAbstract.GetSpider("nowe-technologie", techPage["Spider"])
}

func techComputerWorld() []messages.Message{
	techPage["ComputerWorld"]++
	doc := abstract.GetDoc(fmt.Sprintf("https://www.computerworld.pl/news/archiwum-%v.html", techPage["ComputerWorld"]))

	var news []messages.Message

	abstract.GetDiv(doc,"div.row-list-item").Each(func(i int, s *goquery.Selection){

		image,_ := s.Find("div.row-item-icon > a > figure.frame-responsive img.img-responsive").Attr("src")
		text := s.Find("div.col-lg-9 > a > span.title").Text()
		textlink,_ := s.Find("div.row-item-icon > a").Attr("href")

		temp := messages.Message{
			TitleLink:  textlink,
			Img: messages.Image{
				Header: text,
				ImageUrl: image,
			},
		}

		if !temp.Img.IsEmpty() && temp.TitleLink != ""{
			news = append(news,temp)
		}
	})
	return news
}