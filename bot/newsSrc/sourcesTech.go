package newsSrc

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
	"bot-git/bot/messages"
	"bot-git/bot/newsSrc/newsAbstract"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var GetTech = []newsAbstract.GetNews{
	techSpider,
	techComputerWorld,
	techWirtualneMedia,
}

var TechPage = map[string]int{
	"Spider": 0,
	"ComputerWorld": 0,
	"WirtualneMedia":0,
}

func techSpider() []messages.Message{
	blacklists.New("techSpiderBL")
	TechPage["Spider"]++
	return newsAbstract.GetSpider("nowe-technologie", TechPage["Spider"])
}

func techWirtualneMedia() []messages.Message{
	blacklists.New("techWirtualneMediaBL")
	TechPage["WirtualneMedia"]++
	return newsAbstract.GetWirtualneMedia("technologie", TechPage["WirtualneMedia"])
}

func techComputerWorld() []messages.Message{
	blacklists.New("techComputerWorldBL")
	TechPage["ComputerWorld"]++
	doc := abstract.GetDoc(fmt.Sprintf("https://www.computerworld.pl/news/archiwum-%v.html", TechPage["ComputerWorld"]))

	var news []messages.Message

	abstract.GetDiv(doc,"div.row-list-item").Each(func(i int, s *goquery.Selection){

		image,_ := s.Find("div.row-item-icon > a > figure.frame-responsive img.img-responsive").Attr("src")
		text := s.Find("div.col-lg-9 > a > span.title").Text()
		textlink,_ := s.Find("div.row-item-icon > a").Attr("href")

		temp := messages.Message{
			TitleLink:  fmt.Sprintf("https://www.computerworld.pl/%v", textlink),
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