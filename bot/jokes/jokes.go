package jokes

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
	"bot-git/bot/limit"
	"bot-git/config"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type getJoke func() []string

// all getJoke functions' algorithm:
// 1. open the joke website
// 2. get the joke's div
// 3. get rid of unnecessary text and whitespace

var jokeList []string

func Fetch(hard bool) string {
	limit.AddRequest(abstract.GetUserId(), "joke")
	var jokeSources []getJoke

	if hard {
		jokeSources = jokersHard
	} else {
		jokeSources = jokersPl
		if checkDay() {
			jokeSources = jokersEn
		}
	}
	var jokeFunction getJoke
	canReturn := false
	var joke string
	for canReturn == false {
		if len(jokeList) == 0 {
			jokeFunction = jokeSources[rand.Intn(len(jokeSources))]
			jokeList = jokeFunction()
		}
		joke = getRandomJoke(jokeList)
		canReturn = handleBlacklist(jokeFunction, joke)
	}
	return joke
}

func checkDay() bool {
	return time.Now().Weekday().String() == config.BotCfg.EnglishDay
}

func getFunctionName(functionReturningJoke getJoke) string {
	return runtime.FuncForPC(reflect.ValueOf(functionReturningJoke).Pointer()).Name()
}

func getRandomJoke(jokeList []string) string {
	return jokeList[rand.Intn(len(jokeList))]
}

func handleBlacklist(functionReturningJoke getJoke, jokeReturned string) bool {
	blacklist := blacklists.BlacklistsMap[fmt.Sprintf("%vBL", getFunctionName(functionReturningJoke))]

	if blacklist.Contains(jokeReturned) {
		removeFromJokeList(jokeReturned)
		return false
	}

	blacklist.AddElement(jokeReturned)
	return true
}

func removeFromJokeList(joke string) {
	for i, v := range jokeList {
		if v == joke {
			jokeList[i] = jokeList[len(jokeList)-1]
			jokeList = jokeList[:len(jokeList)-1]
			return
		}
	}
}

func getJokesList(selectionsToFormat *goquery.Selection) []string {

	var jokes []string
	selectionsToFormat.Each(func(i int, s *goquery.Selection) {
		selectionHTML, _ := s.Html()
		jokes = append(jokes, fixFormat(selectionHTML))
	})

	return jokes
}

func fixFormat(HTMLtoFormat string) string {
	formattedString := strings.ReplaceAll(HTMLtoFormat, "<br>", "\n")
	formattedString = strings.ReplaceAll(formattedString, "<br/>", "\n")
	formattedString = strings.ReplaceAll(formattedString, "<p>", "")
	formattedString = strings.ReplaceAll(formattedString, "</p>", "")

	// markdown escape
	formattedString = strings.ReplaceAll(formattedString, "-", "\\-")
	formattedString = strings.ReplaceAll(formattedString, "*", "\\*")

	formattedString = strings.TrimSpace(formattedString)

	return formattedString
}
