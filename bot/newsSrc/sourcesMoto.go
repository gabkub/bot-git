package newsSrc

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc/newsAbstract"
)

var GetMoto = []newsAbstract.GetNews{
	motoAutoCentrum,
}

var motoPage = map[string]int{
	"AutoCentrum": 0,
	"Moto": 0,
}

func motoAutoCentrum() []messages.Message{
	blacklists.New("motoAutoCentrumBL")
	motoPage["AutoCentrum"]++
	doc := abstract.GetDoc(fmt.Sprintf("https://www.autocentrum.pl/newsy/strona-%v/", motoPage["AutoCentrum"]))
	var news []messages.Message

	abstract.GetDiv(doc,"div.ac-article-wrapper").Each(func(i int, s *goquery.Selection){

		image,_ := s.Find("a > div.photo > picture > img.img-responsive").Attr("src")
		text,_ := s.Find("a > div.photo > picture > img.img-responsive").Attr("alt")
		textlink,_ := s.Find("a.news-box").Attr("href")
		temp := messages.Message{
			TitleLink:  fmt.Sprintf("https://www.autocentrum.pl%v",textlink),
			Img: messages.Image{
				Header: text,
				ImageUrl: fmt.Sprintf("https://www.autocentrum.pl%v", image),
			},
		}

		if !temp.Img.IsEmpty() && temp.TitleLink != ""{
			news = append(news,temp)
		}
	})

	return news
}
func motoMoto() []messages.Message{
	motoPage["Moto"]++
	doc := abstract.GetDoc(fmt.Sprintf("http://moto.pl/MotoPL/0,88389.html?str=%v_24561775", motoPage["Moto"]))

	div := abstract.GetDiv(doc,"li.entry")

	var news []messages.Message

	div.Each(func(i int, s *goquery.Selection){

		image,_ := s.Find("a > fiugre.imgw > img").Attr("src")
		text,_ := s.Find("a").Attr("title")
		textlink,_ := s.Find("a").Attr("href")
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