package memes

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklist"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"math/rand"
	"reflect"
	"runtime"
)

type getMeme func() []config.Image

var MemeList []config.Image

func Fetch() config.Image {
	var memeF getMeme
	if len(MemeList)==0 {
		memeF = memSources[rand.Intn(len(memSources))]
		MemeList = memeF()
	}

	m := MemeList[rand.Intn(len(MemeList))]
	handleBL(memeF, m.ImageUrl)

	return m
}

func getFunctionName(f getMeme) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func handleBL(f getMeme, meme string) {
	bl := blacklist.MapBL[getFunctionName(f)]

	if bl.Contains(meme) {
		return
	}

	bl.Add(meme)

	for i,v := range MemeList {
		if v.ImageUrl == meme {
			MemeList[i] = MemeList[len(MemeList)-1]
			MemeList = MemeList[:len(MemeList)-1]
			return
		}
	}
}