package abstract

import (
	"bot-git/bot/messages"
	"bot-git/logg"
	"github.com/PuerkitoBio/goquery"
	"github.com/mattermost/mattermost-server/model"
	"log"
	"net/http"
	"strings"
)

var MsgChannel *model.Channel

type ReactForMsgs []string

func (r ReactForMsgs) ContainsMessage(msg string) bool {
	for _, v := range r {
		if strings.Contains(msg, v) {
			return true
		}
	}
	return false
}

type Handler interface {
	CanHandle(msg string) bool
	Handle(msg string) messages.Message
	GetHelp() messages.Message
}

// TODO this is not the place for this, refactor !!!
func GetDoc(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logg.WriteToFile("Error opening the joke/meme website.")
		log.Fatal("Error while opening the website. Error: " + err.Error())
	}
	return doc
}

// TODO this is not the place for this, refactor !!!
func GetDiv(d *goquery.Document, container string) *goquery.Selection {
	// get the random joke website shows
	div := d.Find(container)
	if div == nil {
		logg.WriteToFile("Error scraping the jokes/memes.")
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
