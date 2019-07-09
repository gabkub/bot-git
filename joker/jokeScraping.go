package joker

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// TODO
// REFACTOR
// which div contains a joke on a site
var jokeDiv = map[string][]string{
	"https://icanhazdadjoke.com/": []string{"div.card-content", ""},
	"https://perelki.net/random": []string{"div.container:first-child", "div.about"},
}

func GetBody(url string) (string, error){

	doc, err := goquery.NewDocument(url)

	if err != nil{
		return "", err
	}


	joke := FindRandom(doc,jokeDiv[url][0],jokeDiv[url][1])

	return joke, nil
}

func FindRandom(doc *goquery.Document, tofind, except string) string{
	div := doc.Find(tofind)

	if div == nil{
		return ""
	}

	result := div.Text()

	if except != ""{
		result = strings.ReplaceAll(div.Text(), doc.Find(except).Text(), "")
	}

	result = strings.TrimSpace(result)

	return result
}



