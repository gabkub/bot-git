package commands

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/blacklists"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/newsSrc/newsAbstract"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
)

type news struct {
	commands []string
}

var NewsHandler news
var resultsNews = make(map[string][]messages.Message)

func (n *news) New() abstract.Handler {
	n.commands = []string{"news"}
	return n
}

func (n *news) CanHandle(msg string) bool {
	return abstract.FindCommand(n.commands, msg)
}

func (n *news) Handle(msg string) messages.Message {
	if strings.Contains(msg, "-h") {
		return n.GetHelp()
	}
	msgSplit := strings.Split(msg," ")
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
	var news messages.Message
	for canReturn==false {
		if len(resultsNews[functionName]) == 0 {
			resultsNews[functionName] = newsFunction()
		}
		news = getRandomNews(functionName)
		canReturn =	handleBlacklist(newsFunction, news)
	}
	messages.Response = news

	return messages.Response
}
func getRandomNews(mapName string) messages.Message{
	result := resultsNews[mapName][rand.Intn(len(resultsNews[mapName]))]
	removeFromNewsList(result)

	return result
}

func getFunctionName(functionReturningNews newsAbstract.GetNews) string {
	return runtime.FuncForPC(reflect.ValueOf(functionReturningNews).Pointer()).Name()
}

func handleBlacklist(functionReturningNews newsAbstract.GetNews, newsReturned messages.Message) bool {
	blacklist := blacklists.BlacklistsMap[fmt.Sprintf("%vBL", getFunctionName(functionReturningNews))]

	if blacklist.Contains(newsReturned.TitleLink) {
		removeFromNewsList(newsReturned)
		return false
	}

	blacklist.AddElement(newsReturned.TitleLink)
	return true
}


func removeFromNewsList(news messages.Message) {
	for k,v := range resultsNews {
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
func (n *news) GetHelp() messages.Message {
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
	messages.Response.Text = sb.String()
	return messages.Response
}
