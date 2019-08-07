package newsSrc

import (
	"bot-git/bot/blacklists"
	"bot-git/bot/newsSrc/newsAbstract"
)

var GetMedia = []newsAbstract.GetNews{
	mediaWirtualneMedia,
	mediaSpider,
}

var MediaPage = map[string]int{
	"Spider":         0,
	"WirtualneMedia": 0,
}

func mediaSpider() []*newsAbstract.News {
	blacklists.New("mediaSpiderBL")
	MediaPage["Spider"]++
	return getSpider("media", MediaPage["Spider"])
}

func mediaWirtualneMedia() []*newsAbstract.News {
	blacklists.New("mediaWirtualneMediaBL")
	MediaPage["WirtualneMedia"]++
	return getWirtualneMedia("kultura-i-rozrywka", MediaPage["WirtualneMedia"])
}
