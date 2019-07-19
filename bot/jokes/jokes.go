package jokes

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type getJoke func() string

// all getJoke functions' algorithm:
// 1. open the joke website
// 2. get the joke's div
// 3. get rid of unnecessary text and whitespace

func Fetch() string {
	limit.AddRequest(abstract.GetUserId(), "joke")
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

func getFunctionName(functionReturningJoke getJoke) string {
	return runtime.FuncForPC(reflect.ValueOf(functionReturningJoke).Pointer()).Name()
}

func handleBlacklist(functionReturningJoke getJoke, joke string) {
	blacklist := blacklists.MapBL[getFunctionName(functionReturningJoke)]

	if blacklist.Contains(joke) {
		functionReturningJoke()
	}

	blacklist.Add(joke)
}

func fixFormat(toFormatHTML string) string {

	formattedString := strings.ReplaceAll(toFormatHTML, "<br>", "\n")
	formattedString = strings.ReplaceAll(formattedString, "<br/>", "\n")

	// markdown escape
	formattedString = strings.ReplaceAll(formattedString, "-", "\\-")
	formattedString = strings.ReplaceAll(formattedString, "*", "\\*")

	formattedString = strings.TrimSpace(formattedString)

	return formattedString
}