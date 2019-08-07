package memes

import (
	"bot-git/bot/blacklists"
	"bot-git/bot/messages"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var memSources = []getMeme{
	memedroid,
}

var countMemedroid = 1

func memedroid() []messages.Image {

	blacklists.New("getMemedroidBL")
	var memes []messages.Image

	div := contentFetcher.Fetch(fmt.Sprintf("https://www.memedroid.com/memes/top/week/%v", countMemedroid), "article.gallery-item")
	div.Each(func(i int, s *goquery.Selection) {

		image, _ := s.Find("a.dyn-link:nth-child(2) img.img-responsive").Attr("src")
		temp := messages.Image{
			Header:   s.Find("header.item-header h1").Text(),
			ImageUrl: image,
		}

		if temp.ImageUrl != "" {
			memes = append(memes, temp)
		}
	})
	countMemedroid++
	return memes
}
