package meme

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type meme struct{
	header 		string
	imageUrl 	string
	lastMeme	int
}

func (result *meme) getMemedroid(url string) error {
	doc, err := goquery.NewDocument(url)

	var helper  = helpFindMeme{"article.gallery-item:nth-child(%d) div.item-aux-container", 1, "header.item-header h1",
		"a.dyn-link:nth-child(2)"}

	if err != nil{
		return  err
	}

	helper.randContainer()

	result.header, result.imageUrl, result.lastMeme = findCommand(doc, helper)

	for result.isEmpty() {
		helper.randContainer()
		result.header, result.imageUrl, result.lastMeme = findCommand(doc, helper)
	}

	return  nil
}

func (result meme) isEmpty() bool{

	if result.lastMeme == 0 || result.header == "" || result.imageUrl == ""{
		return true
	}

	return false
}
func findCommand(doc *goquery.Document, helper helpFindMeme) (string,string,int){
	div := doc.Find(helper.mainContainer)

	if div == nil{
		return "","",0
	}
	image := div.Find(helper.image)
	header := div.Find(helper.header)

	//Get html node
	imageNode,_ := image.Html()
	imageUrl := getImageUrlFromNode(imageNode)


	if header != nil && imageUrl != ""{
		return header.Text(),
			imageUrl,
			helper.mainContainerID
	}

	return "","",0
}

func getImageUrlFromNode(imageHTML string) string {

	if strings.HasPrefix(imageHTML, "<img src=\"") {
		imageHTML = strings.ReplaceAll(imageHTML, "<img src=\"", "")

		for i := 0; i < len(imageHTML); i++ {

			if string(imageHTML[i]) == "\"" {
				imageHTML = imageHTML[:i]

				return imageHTML
			}
		}
	}

	return ""

}