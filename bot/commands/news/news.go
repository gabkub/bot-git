package news

import (
	"bot-git/bot/abstract"
	"bot-git/bot/newsSrc"
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/messageBuilders"
	"fmt"
	"math/rand"
	"strings"
)

type news struct {
}

var commands abstract.ReactForMsgs = []string{"news"}

func New() *news {
	return &news{}
}

func (n *news) CanHandle(msg string) bool {
	return commands.ContainsMessage(msg)
}

var newsFeedsMap = map[string][]newsAbstract.GetNews{
	"games":   newsSrc.GetGame,
	"gry":     newsSrc.GetGame,
	"media":   newsSrc.GetMedia,
	"science": newsSrc.GetScience,
	"nauka":   newsSrc.GetScience,
	"tech":    newsSrc.GetTech,
	"news":    newsSrc.GetTech,
	"travel":  newsSrc.GetVoyage,
	"podróże": newsSrc.GetVoyage,
	"moto":    newsSrc.GetMoto,
}

func (n *news) Handle(msg string, sender abstract.MessageSender) {
	msgSplit := strings.Split(msg, " ")
	cat := msgSplit[len(msgSplit)-1]

	news := getFreshNews(cat)

	sender.Send(messageBuilders.News(news))
}

func getFreshNews(cat string) *newsAbstract.News {
	newsFunctions, ok := newsFeedsMap[cat]
	if !ok {
		newsFunctions = newsSrc.GetTech
	}

	tmp := make([]newsAbstract.GetNews, len(newsFunctions))
	copy(tmp, newsFunctions)

	for len(tmp) > 0 {
		i := rand.Intn(len(tmp))
		getNews := tmp[i]
		n, ok := getNews()
		if ok {
			return n
		}
		tmp = append(tmp[:i], tmp[i+1:]...)
	}
	return notFoundNews(cat)
}

func notFoundNews(cat string) *newsAbstract.News {
	text := fmt.Sprintf("Nie mam nic nowego w katergori: _%s_", cat)
	img := abstract.NewImage(text, "https://png.pngtree.com/svg/20170217/7c2ce8d09c.svg")
	return newsAbstract.NewNews(text, img)
}
