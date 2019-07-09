package joker

import (
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"math/rand"
	"time"
)

var eng = []string{"https://icanhazdadjoke.com/"}
var pl = []string{"https://perelki.net/random"}

//type getJoker func() string
//
//var jokerPl []getJoker
//var jokerEn []getJoker
//
//type Joker interface {
//	GetJoke()
//}

func Fetch(cfg *config.BotConfig) (string, error) {
	//var j []Joker
	//if isEnglish {
	//	j = jokerEn
	//}
	//
	// jj := j[randomindex]
	//return jj.GetJoke()
	url := getUrl(checkDay(cfg))

	joke, err := GetBody(url)

	if err != nil {
		return "", err
	}

	return joke, nil
}

func checkDay(cfg *config.BotConfig) bool {

	if time.Now().Weekday().String() == cfg.EnglishDay {
		return true
	}
	return false
}

func getUrl(isEng bool) string {

	if isEng {
		return eng[rand.Intn(len(eng))]
	}

	return pl[rand.Intn(len(pl))]
}
