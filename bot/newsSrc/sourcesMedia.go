package newsSrc

import (
	"bot-git/bot/newsSrc/newsAbstract"
)

var GetMedia = []newsAbstract.GetNews{
	mediaWirtualneMedia,
	mediaSpider,
}

func mediaSpider() (*newsAbstract.News, bool) {
	return getSpider("media")
}

func mediaWirtualneMedia() (*newsAbstract.News, bool) {
	return getWirtualneMedia("kultura-i-rozrywka")
}
