package memes

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklist"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var memSources = []getMeme{
	getMemedroid,
}

var countMemedroid = 1

func getMemedroid() []config.Image {
	countMemedroid++
	blacklist.New("getMemedroidBL")

	doc := abstract.GetDoc(fmt.Sprintf("https://www.memedroid.com/memes/top/day/%v", countMemedroid))
	div := abstract.GetDiv(doc, "article.gallery-item")

	var memes []config.Image

	div.Each(func(i int, s *goquery.Selection){

		image,_ := s.Find("a.dyn-link:nth-child(2) img.img-responsive").Attr("src")
		temp := config.Image{
			s.Find("header.item-header h1").Text(),
			image,
		}

		if !temp.IsEmpty() {
			memes = append(memes,temp)
		}
	})
	return memes
}
