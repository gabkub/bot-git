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

var MemeList []messages.Image

func Fetch() messages.Image {
	limit.AddRequest(abstract.GetUserId(), "meme")
	var memeFunction getMeme
	if len(MemeList)==0 {
		memeFunction = memSources[rand.Intn(len(memSources))]
		MemeList = memeFunction()
	}

	meme := MemeList[rand.Intn(len(MemeList))]
	handleBL(memeFunction, meme.ImageUrl)

	return meme
}

func getFunctionName(functionReturningMeme getMeme) string {
	return runtime.FuncForPC(reflect.ValueOf(functionReturningMeme).Pointer()).Name()
}

func handleBL(functionReturningMeme getMeme, memeReturned string) {
	bl := blacklists.MapBL[getFunctionName(functionReturningMeme)]

	if bl.Contains(memeReturned) {
		return
	}

	bl.Add(memeReturned)

	for i,v := range MemeList {
		if v.ImageUrl == memeReturned {
			MemeList[i] = MemeList[len(MemeList)-1]
			MemeList = MemeList[:len(MemeList)-1]
			return
		}
	}
}