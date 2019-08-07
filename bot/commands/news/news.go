package news

import (
	"bot-git/bot/abstract"
	"bot-git/bot/blacklists"
	"bot-git/bot/newsSrc"
	"bot-git/bot/newsSrc/newsAbstract"
	"bot-git/messageBuilders"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
)

type news struct {
	commands abstract.ReactForMsgs
}

var resultsNews = make(map[string][]*newsAbstract.News)

func New() *news {
	return &news{[]string{"news"}}
}

func (n *news) CanHandle(msg string) bool {
	return n.commands.ContainsMessage(msg)
}

func (n *news) Handle(msg string, sender abstract.MessageSender) {
	if strings.Contains(msg, "-h") {
		sender.Send(messageBuilders.Text(n.GetHelp()))
		return
	}
	msgSplit := strings.Split(msg, " ")
	cat := msgSplit[len(msgSplit)-1]
	var newsFunction newsAbstract.GetNews
	switch cat {
	case "games":
		newsFunction = newsSrc.GetGame[rand.Intn(len(newsSrc.GetGame))]
	case "gry":
		newsFunction = newsSrc.GetGame[rand.Intn(len(newsSrc.GetGame))]
	case "media":
		newsFunction = newsSrc.GetMedia[rand.Intn(len(newsSrc.GetMedia))]
	case "science":
		newsFunction = newsSrc.GetScience[rand.Intn(len(newsSrc.GetScience))]
	case "nauka":
		newsFunction = newsSrc.GetScience[rand.Intn(len(newsSrc.GetScience))]
	case "tech":
		newsFunction = newsSrc.GetTech[rand.Intn(len(newsSrc.GetTech))]
	case "news":
		newsFunction = newsSrc.GetTech[rand.Intn(len(newsSrc.GetTech))]
	case "travel":
		newsFunction = newsSrc.GetVoyage[rand.Intn(len(newsSrc.GetVoyage))]
	case "podróże":
		newsFunction = newsSrc.GetVoyage[rand.Intn(len(newsSrc.GetVoyage))]
	case "moto":
		newsFunction = newsSrc.GetMoto[(rand.Intn(len(newsSrc.GetMoto)))]
	}

	if newsFunction == nil {
		newsFunction = newsSrc.GetTech[rand.Intn(len(newsSrc.GetTech))]
	}

	functionName := getFunctionName(newsFunction)
	canReturn := false
	var news *newsAbstract.News
	for canReturn == false {
		if len(resultsNews[functionName]) == 0 {
			resultsNews[functionName] = newsFunction()
		}
		news = getRandomNews(functionName)
		canReturn = handleBlacklist(newsFunction, news)
	}
	sender.Send(messageBuilders.News(news))
}
func getRandomNews(mapName string) *newsAbstract.News {
	result := resultsNews[mapName][rand.Intn(len(resultsNews[mapName]))]
	removeFromNewsList(result)

	return result
}

func getFunctionName(functionReturningNews newsAbstract.GetNews) string {
	return runtime.FuncForPC(reflect.ValueOf(functionReturningNews).Pointer()).Name()
}

func handleBlacklist(functionReturningNews newsAbstract.GetNews, newsReturned *newsAbstract.News) bool {
	blacklist := blacklists.BlacklistsMap[fmt.Sprintf("%vBL", getFunctionName(functionReturningNews))]

	if blacklist.Contains(newsReturned.TitleLink) {
		removeFromNewsList(newsReturned)
		return false
	}

	blacklist.AddElement(newsReturned.TitleLink)
	return true
}

func removeFromNewsList(news *newsAbstract.News) {
	for k, v := range resultsNews {
		for i, value := range v {
			if value.TitleLink == news.TitleLink {
				v[i] = v[len(v)-1]
				v = v[:len(v)-1]
				resultsNews[k] = v
				return
			}
		}
	}
}
func (n *news) GetHelp() string {
	var sb strings.Builder
	sb.WriteString("Losowy news.\n\n")
	sb.WriteString("Dostępne kategorie:\n")
	sb.WriteString("- gry/games\n")
	sb.WriteString("- media\n")
	sb.WriteString("- nauka/science\n")
	sb.WriteString("- tech (domyślna)\n")
	sb.WriteString("- moto\n")
	sb.WriteString("- podróże/travel\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_news_\n")
	return sb.String()
}
