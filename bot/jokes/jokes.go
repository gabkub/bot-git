package jokes

import (
	"bot-git/bot/abstract"
	"bot-git/bot/limit"
	"bot-git/config"
	"math/rand"
	"time"
)

type getJoke func() (*string, bool)

func Fetch(userId abstract.UserId, hard bool) string {
	limit.AddJoke(userId)
	var jokeSources []getJoke

	if hard {
		jokeSources = jokersHard
	} else {
		jokeSources = jokersPl
		if isEnglishDay() {
			jokeSources = jokersEn
		}
	}

	tmp := make([]getJoke, len(jokeSources))
	copy(tmp, jokeSources)

	for len(tmp) > 0 {
		i := rand.Intn(len(tmp))
		getJoke := tmp[i]
		joke, ok := getJoke()
		if ok {
			return *joke
		}
		tmp = append(tmp[:i], tmp[i+1:]...)
	}
	return notFoundText()
}

func notFoundText() string {
	if isEnglishDay() {
		return "Nothing new"
	}
	return "Nie mam nic nowego"
}

func isEnglishDay() bool {
	return time.Now().Weekday().String() == config.BotCfg.EnglishDay
}
