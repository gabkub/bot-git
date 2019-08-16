package memes

import (
	"bot-git/bot/abstract"
	"bot-git/contentFetcher"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var memSources = []getMeme{
	memedroid,
}

func memedroid() (*abstract.Image, bool) {
	for i := 1; i <= 3; i++ {
		meme, ok := fetchMeme(i)
		if ok {
			return meme, true
		}
	}
	return nil, false
}

func fetchMeme(page int) (*abstract.Image, bool) {
	var meme *abstract.Image

	div := contentFetcher.Fetch(fmt.Sprintf("https://www.memedroid.com/memes/top/week/%v", page), "article.gallery-item")
	div.Each(func(i int, s *goquery.Selection) {
		if meme != nil {
			return
		}
		image, _ := s.Find("a.dyn-link:nth-child(2) img.img-responsive").Attr("src")
		temp := &abstract.Image{
			Header:   s.Find("header.item-header h1").Text(),
			ImageUrl: image,
		}
		if temp.ImageUrl != "" && Blacklist.IsFresh(&temp.Header) {
			meme = temp
		}
	})

	return meme, meme != nil
}
