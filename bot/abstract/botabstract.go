package abstract

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/logs"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

var limitMessages = []string{
	"Do roboty!", "Hej ho, hej ho, do pracy by się szło...", "Już się zmęczyłem.", "Zostaw mnie w spokoju.",
	"Koniec śmieszków...", "Foch.", "Nie.", "Zaraz wracam. Albo i nie...", "A może by tak popracować?", "~~żart~~",
}

func RandomLimitMsg() messages.Message {
	var msg messages.Message
	msg.New()
	msg.Text = limitMessages[rand.Intn(len(limitMessages))]
	return msg
}

type Handler interface {
	CanHandle(msg string) bool
	Handle(msg string) messages.Message
	GetHelp() messages.Message
}

func FindCommand(commands []string, msg string) bool {
	for _,v := range commands {
		if strings.Contains(msg, v){
			return true
		}
	}
	return false
}

func GetDoc(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil{
		logs.WriteToFile("Error opening the joke/meme website.")
		log.Fatal("Error while opening the website. Error: " + err.Error())
	}
	return doc
}

func GetDiv(d *goquery.Document, container string) *goquery.Selection {
	// get the random joke website shows
	div := d.Find(container)
	if div == nil{
		logs.WriteToFile("Error scraping the jokes/memes.")
		log.Fatal("Error scraping the jokes/memes.")

	}
	return div
}

var userId string

func GetUserId() string {
	return userId
}

func SetUserId(id string) {
	userId = id
}
