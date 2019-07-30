package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
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
var resultsNews = map[string][]messages.Message{}

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
	case "voyage":
		newsFunction = newsSrc.GetVoyage[rand.Intn(len(newsSrc.GetVoyage))]
	case "podróże":
		newsFunction = newsSrc.GetVoyage[rand.Intn(len(newsSrc.GetVoyage))]
	}

	if newsFunction != nil {
		functionName := getFunctionName(newsFunction)
		if len(resultsNews[functionName]) == 0 {
			resultsNews[functionName] = newsFunction()
	}
		messages.Response = randomNews(functionName)
	} else {
		messages.Response.Text = "Brak kategorii. Sprawdź listę dostępnych kategorii komendą _news -h_"
	}

	return messages.Response
}
func randomNews(mapName string) messages.Message{
	result := resultsNews[mapName][rand.Intn(len(resultsNews[mapName]))]
	removeFromNewsList(result)

	return result
}
func getFunctionName(functionReturningNews newsAbstract.GetNews) string {
	return runtime.FuncForPC(reflect.ValueOf(functionReturningNews).Pointer()).Name()
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
	sb.WriteString("Losowy news z polskich i zagranicznych witryn.\n\n")
	sb.WriteString("Dostępne kategorie:\n")
	sb.WriteString("- games/gry\n")
	sb.WriteString("- media\n")
	sb.WriteString("- science/nauka\n")
	sb.WriteString("- tech\n")
	sb.WriteString("- voyage/podróże\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_news_\n")
	messages.Response.Text = sb.String()
	return messages.Response
}
