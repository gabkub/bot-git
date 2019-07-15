package abstract

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"log"
	"net/http"
	"os"
	"strings"
)


type Handler interface {
	CanHandle(msg string) bool
	Handle(msg string) config.Msg
	GetHelp() config.Msg
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
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil{
		log.Fatal("Error while opening the website. Error: " + err.Error())
		os.Exit(1)
	}

	resp.Body.Close()
	return doc
}

func GetDiv(d *goquery.Document, container string) *goquery.Selection {
	// get the random joke website shows
	div := d.Find(container)
	if div == nil{
		log.Fatal("Empty joke.")
		os.Exit(1)
	}
	return div
}