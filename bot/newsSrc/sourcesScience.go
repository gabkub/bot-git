package newsSrc

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc/newsAbstract"
)

var GetScience = []newsAbstract.GetNews{
	sciencePrzystanek,
	scienceSpider,
}

var sciencePage = map[string]int{
	"Spider": 0,
	"Przystanek": -1,
}
func scienceSpider() []messages.Message{
	blacklists.New("scienceSpiderBL")
	sciencePage["Spider"]++
	return newsAbstract.GetSpider("nauka", sciencePage["Spider"])
}

func sciencePrzystanek() []messages.Message{
	blacklists.New("sciencePrzystanekBL")
	sciencePage["Przystanek"]++
	doc := abstract.GetDoc(fmt.Sprintf("http://przystaneknauka.us.edu.pl/news?page=%v", sciencePage["Przystanek"]))

	div := abstract.GetDiv(doc,"div.views-row")

	var news []messages.Message

	div.Each(func(i int, s *goquery.Selection){

		image,_ := s.Find("a > img").Attr("src")
		text := s.Find("h3.title").Text()
		textlink, _ := s.Find("h3.title > a").Attr("href")
		temp := messages.Message{
			TitleLink:  fmt.Sprintf("http://przystaneknauka.us.edu.pl%v",textlink),
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
