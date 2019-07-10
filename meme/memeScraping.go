package meme

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"sync"
)

type meme struct{
	header 		string
	imageUrl 	string
	blackList	[]string
}


func (result *meme) getMemedroid(url string) error {
	doc, err := goquery.NewDocument(url)

	var helper  = helpFindMeme{"article.gallery-item:nth-child(%d) div.item-aux-container", 1, "header.item-header h1",
		"a.dyn-link:nth-child(2)" }

	if err != nil{
		return  err
	}

	result.random(doc, &helper)


	return  nil
}

func (result meme) isEmpty() bool{

	if result.header == "" || result.imageUrl == ""{
		return true
	}

	return false
}

func (result *meme) random(document *goquery.Document, help *helpFindMeme) {

	help.randContainer()
	result.findCommand(document, help)

	//println(result.header)
	//println(result.imageUrl)
	if result.isEmpty() {
		result.random(document,help)
	}


}

func (result *meme) isOnBlackList(param string) bool {

	if param == ""{
		return true
	}

	var m  = &sync.Mutex{}
	m.Lock()
	defer m.Unlock()
	for _, element := range result.blackList{
		if param == element{
			return true
		}
	}
	
	return false
}
func (result *meme)findCommand(doc *goquery.Document, helper *helpFindMeme){
	div := doc.Find(helper.mainContainer)

	if div == nil{
		result = &meme{}
		return
	}
	image := div.Find(helper.image)
	header := div.Find(helper.header)

	//Get html node
	imageNode,_ := image.Html()
	imageUrl := getImageUrlFromNode(imageNode)

	if !result.isOnBlackList(imageUrl){
		result.header = header.Text()
		result.imageUrl = imageUrl
		result.blackList = append(result.blackList, imageUrl)
		return
	}

	result.random(doc,helper)
	return
}

func getImageUrlFromNode(imageHTML string) string{

	if strings.HasPrefix(imageHTML,"<img src=\""){
		imageHTML = strings.ReplaceAll(imageHTML,"<img src=\"","")

		for i := 0; i < len(imageHTML); i++{

			if string(imageHTML[i]) == "\""{
				imageHTML = imageHTML[:i]

				return imageHTML
			}
		}
	}

	return ""
}