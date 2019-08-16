package memes

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklist"
	"bot-git/bot/limit"
	"bot-git/config"
	"math/rand"
)

var Blacklist = blacklist.New(20)

type getMeme func() (*abstract.Image, bool)

func Fetch(userId abstract.UserId) *abstract.Image {
	limit.AddMeme(userId)

	tmp := make([]getMeme, len(memSources))
	copy(tmp, memSources)

	for len(tmp) > 0 {
		i := rand.Intn(len(tmp))
		getMeme := tmp[i]
		meme, ok := getMeme()
		if ok {
			return meme
		}
		tmp = append(tmp[:i], tmp[i+1:]...)
	}
	return notFoundImage()
}

func notFoundImage() *abstract.Image {
	title := "Nic nowego"
	if config.IsEnglishDay() {
		title = "Nothing new"
	}
	return &abstract.Image{
		Header:   title,
		ImageUrl: config.NothingNewImageUrl,
	}
}
