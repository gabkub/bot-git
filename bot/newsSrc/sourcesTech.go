package newsSrc

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var GetTech = []newsAbstract.GetNews{
	techSpider,
	techComputerWorld,
	techWirtualneMedia,
}

var TechPage = map[string]int{
	"Spider":         0,
	"ComputerWorld":  0,
	"WirtualneMedia": 0,
}

func techSpider() []*newsAbstract.News {
	blacklists.New("techSpiderBL")
	TechPage["Spider"]++
	return getSpider("nowe-technologie", TechPage["Spider"])
}

func techWirtualneMedia() []*newsAbstract.News {
	blacklists.New("techWirtualneMediaBL")
	TechPage["WirtualneMedia"]++
	return getWirtualneMedia("technologie", TechPage["WirtualneMedia"])
}

func techComputerWorld() []*newsAbstract.News {
	blacklists.New("techComputerWorldBL")
	TechPage["ComputerWorld"]++
	var news []*newsAbstract.News
	div := contentFetcher.Fetch(fmt.Sprintf("https://www.computerworld.pl/news/archiwum-%v.html", TechPage["ComputerWorld"]), "div.row-list-item")
	div.Each(func(i int, s *goquery.Selection) {
		image, _ := s.Find("div.row-item-icon > a > figure.frame-responsive img.img-responsive").Attr("src")
		text := s.Find("div.col-lg-9 > a > span.title").Text()
		textLink, _ := s.Find("div.row-item-icon > a").Attr("href")
		link := fmt.Sprintf("https://www.computerworld.pl/%v", textLink)
		img := abstract.NewImage(text, image)
		temp := newsAbstract.NewNews(link, img)
		if !temp.Img.IsEmpty() && temp.TitleLink != "" {
			news = append(news, temp)
		}
	})
	return news
}
