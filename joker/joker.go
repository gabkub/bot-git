package joker

import (
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"math/rand"
	"time"
)

var eng = []string{"https://icanhazdadjoke.com/"}
var pl = []string{"https://perelki.net/random"}

func Fetch(cfg *config.BotConfig) (string,error){

	url := GetUrl(CheckWednesday(cfg))

	joke, err := GetBody(url)

	if err != nil{
		return "", err
	}

	return joke, nil
}

func CheckWednesday(cfg *config.BotConfig) bool {

	if time.Now().Weekday().String() == cfg.EnglishDay {
		return true
	}
	return false
}

func GetUrl(isEng bool) string {

	if isEng {
		return eng[rand.Intn(len(eng))]
	}

	return pl[rand.Intn(len(pl))]
}