package memes

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
	"bot-git/bot/limit"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
)

type getMeme func() []abstract.Image

var memeList []abstract.Image

func Fetch() abstract.Image {
	limit.AddRequest(abstract.GetUserId(), "meme")
	var memeFunction getMeme

	canReturn := false
	var meme abstract.Image
	for canReturn == false {
		if len(memeList) == 0 {
			memeFunction = memSources[rand.Intn(len(memSources))]
			memeList = memeFunction()
		}
		meme = getRandomMeme(memeList)
		canReturn = handleBlacklist(memeFunction, meme.ImageUrl)
	}

	return meme
}

func getRandomMeme(memeList []abstract.Image) abstract.Image {
	return memeList[rand.Intn(len(memeList))]
}

func getFunctionName(functionReturningMeme getMeme) string {
	return runtime.FuncForPC(reflect.ValueOf(functionReturningMeme).Pointer()).Name()
}

func handleBlacklist(functionReturningJoke getMeme, jokeReturned string) bool {
	blacklist := blacklists.BlacklistsMap[fmt.Sprintf("%vBL", getFunctionName(functionReturningJoke))]

	if blacklist.Contains(jokeReturned) {
		removeFromMemeList(jokeReturned)
		return false
	}

	blacklist.AddElement(jokeReturned)
	return true
}

func removeFromMemeList(meme string) {
	for i, v := range memeList {
		if v.ImageUrl == meme {
			memeList[i] = memeList[len(memeList)-1]
			memeList = memeList[:len(memeList)-1]
			return
		}
	}
}
