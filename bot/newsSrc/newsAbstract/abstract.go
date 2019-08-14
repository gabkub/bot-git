package newsAbstract

import (
	"bot-git/bot/abstract"
)

type GetNews func() (*News, bool)

type News struct {
	TitleLink string
	Img       *abstract.Image
}

func NewNews(titleLink string, img *abstract.Image) *News {
	return &News{TitleLink: titleLink, Img: img}
}
