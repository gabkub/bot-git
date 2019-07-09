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

func Fetch() (string, error) {
	//var j []Joker
	//if isEnglish {
	//	j = jokerEn
	//}
	//
	// jj := j[randomindex]
	//return jj.GetJoke()
	url := getUrl(checkDay())

	joke, err := GetBody(url)

	if err != nil {
		return "", err
	}

	return joke, nil
}

func checkDay() bool {

	if time.Now().Weekday().String() == config.BotCfg.EnglishDay {
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
