package joker

import (
	"../config"
	"math/rand"
	"time"
)

type getJoke func() string

func Fetch() (string, error) {
	jokers := jokerPl
	if checkDay() {
		jokers = jokerEn
	}
	joke := jokers[rand.Intn(len(jokers))]()
	return joke, nil
}

func checkDay() bool {

	if time.Now().Weekday().String() == config.BotCfg.EnglishDay {
		return true
	}
	return false
}
