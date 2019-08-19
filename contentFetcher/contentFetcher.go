package contentFetcher

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func getDoc(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error while opening the website. Error: " + err.Error())
	}
	return doc
}

func Fetch(url, selector string) *goquery.Selection {
	doc := getDoc(url)
	return getDiv(doc, selector)
}

func getDiv(d *goquery.Document, container string) *goquery.Selection {
	// get the random joke website shows
	div := d.Find(container)
	if div == nil {
		log.Fatal("Error scraping the jokes/memes.")
	}
	return div
}
