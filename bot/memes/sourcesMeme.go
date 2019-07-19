package memes

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
)

var memSources = []getMeme{
	getMemedroid,
}

var countMemedroid = 1

func getMemedroid() []messages.Image {
	countMemedroid++
	blacklists.New("getMemedroidBL")

	doc := abstract.GetDoc(fmt.Sprintf("https://www.memedroid.com/memes/top/day/%v", countMemedroid))
	div := abstract.GetDiv(doc, "article.gallery-item")

	var memes []messages.Image

	div.Each(func(i int, s *goquery.Selection){

		image,_ := s.Find("a.dyn-link:nth-child(2) img.img-responsive").Attr("src")
		temp := messages.Image{
			Header: s.Find("header.item-header h1").Text(),
			ImageUrl: image,
		}

		if temp.ImageUrl != "" {
			memes = append(memes,temp)
		}
	})
	return memes
}
