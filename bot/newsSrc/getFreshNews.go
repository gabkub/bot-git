package newsSrc

import (
	"bot-git/bot/abstract"
	"bot-git/bot/newsSrc/newsAbstract"
	"github.com/PuerkitoBio/goquery"
)

func getFreshNews(fetch func(int) *goquery.Selection, try int, sel *selectors) (*newsAbstract.News, bool) {
	to := try + 1
	var result *newsAbstract.News
	for i := 1; i <= to; i++ {
		div := fetch(i)
		div.Each(func(i int, s *goquery.Selection) {
			if result != nil {
				return
			}
			link := sel.linkBuilder(s)
			if newsBlacklist.IsFresh(&link) {
				image := sel.imageSrcSel(s)
				text := sel.imageTextSel(s)
				img := abstract.NewImage(text, image)
				message := newsAbstract.NewNews(link, img)
				if !message.Img.IsEmpty() && message.TitleLink != "" {
					result = message
				}
			}
		})
	}
	return result, result != nil
}

type selectors struct {
	linkBuilder  selFunc
	imageSrcSel  selFunc
	imageTextSel selFunc
}
type selFunc func(*goquery.Selection) string

func newSelectors(linkBuilder selFunc, imageSrcSel selFunc, imageTextSel selFunc) *selectors {
	return &selectors{linkBuilder: linkBuilder, imageSrcSel: imageSrcSel, imageTextSel: imageTextSel}
}
