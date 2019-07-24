package memes

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"math/rand"
	"reflect"
	"runtime"
)

type getMeme func() []messages.Image

var memeList []messages.Image

func Fetch() messages.Image {
	limit.AddRequest(abstract.GetUserId(), "meme")
	var memeFunction getMeme
	if len(memeList)==0 {
		memeFunction = memSources[rand.Intn(len(memSources))]
		memeList = memeFunction()
	}

	meme := memeList[rand.Intn(len(memeList))]
	handleBlacklist(memeFunction, meme.ImageUrl)

	return meme
}

func getFunctionName(functionReturningMeme getMeme) string {
	return runtime.FuncForPC(reflect.ValueOf(functionReturningMeme).Pointer()).Name()
}

func handleBlacklist(functionReturningMeme getMeme, memeReturned string) {
	bl := blacklists.BlacklistsMap[getFunctionName(functionReturningMeme)]

	if bl.Contains(memeReturned) {
		return
	}

	bl.AddElement(memeReturned)

	removeFromMemeList(memeReturned)
}

func removeFromMemeList(meme string) {
	for i,v := range memeList {
		if v.ImageUrl == meme {
			memeList[i] = memeList[len(memeList)-1]
			memeList = memeList[:len(memeList)-1]
			return
		}
	}
}