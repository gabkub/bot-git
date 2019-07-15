package jokes

import (
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklist"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

type getJoke func() string

// all getJoke functions' algorithm:
// 1. open the joke website
// 2. get the joke's div
// 3. get rid of unnecessary text and whitespace

func Fetch() string {
	jokers := jokerPl
	if checkDay() {
		jokers = jokerEn
	}
	joke := jokers[rand.Intn(len(jokers))]()
	return joke
}

func checkDay() bool {
	return time.Now().Weekday().String() == config.BotCfg.EnglishDay
}

func getFunctionName(f getJoke) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func handleBL(f getJoke, joke string) {
	bl := blacklist.MapBL[getFunctionName(f)]

	if bl.Contains(joke) {
		f()
	}

	bl.Add(joke)
}