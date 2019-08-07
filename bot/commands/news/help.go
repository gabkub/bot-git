package news

import (
	"bot-git/bot/abstract"
)

var short = "_news <kategoria>_ - losowy news z danej kategorii (brak kategorii wysyła newsa technologicznego)"
var long = `Losowy news.
Dostępne kategorie:
- gry/games
- media
- nauka/science
- tech (domyślna)
- moto
- podróże/travel
Pełna lista komend:
_news_ -h
`

var help = abstract.NewHelp(short, long)

func (n *news) GetHelp() *abstract.Help {
	return help
}
