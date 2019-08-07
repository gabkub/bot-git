package newsAbstract

import (
	"bot-git/bot/messages"
)

type GetNews func() []*News

type News struct {
	TitleLink string
	Img       *messages.Image
}

func NewNews(titleLink string, img *messages.Image) *News {
	return &News{TitleLink: titleLink, Img: img}
}
