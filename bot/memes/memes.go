package memes

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklist"
	"bot-git/bot/limit"
	"math/rand"
)

var Blacklist = blacklist.New(20)

type getMeme func() []abstract.Image

var memeList []abstract.Image

func Fetch(userId abstract.UserId) abstract.Image {
	limit.AddMeme(userId)
	var memeFunction getMeme

	canReturn := false
	var meme abstract.Image
	for canReturn == false {
		if len(memeList) == 0 {
			memeFunction = memSources[rand.Intn(len(memSources))]
			memeList = memeFunction()
		}
		meme = getRandomMeme(memeList)
		canReturn = handleBlacklist(&meme.ImageUrl)
	}

	return meme
}

func getRandomMeme(memeList []abstract.Image) abstract.Image {
	return memeList[rand.Intn(len(memeList))]
}

func handleBlacklist(meme *string) bool {
	isFresh := Blacklist.IsFresh(meme)
	if !isFresh {
		removeFromMemeList(meme)
		return false
	}
	return true
}

func removeFromMemeList(meme *string) {
	for i, v := range memeList {
		if v.ImageUrl == *meme {
			memeList[i] = memeList[len(memeList)-1]
			memeList = memeList[:len(memeList)-1]
			return
		}
	}
}
