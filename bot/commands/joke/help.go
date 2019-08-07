package joke

import "bot-git/bot/abstract"

var short = "_joke, żart_ - losowy dowcip"
var long = `Wysyła losowy dowcip. W dzień określony w pliku konfiguracyjnym żarty są w języku angielskim.
Limity:
7:00-8:59 - 3 żarty\
9:00-14:59 - 1 żart na godzinę
15:00-6:59 - brak limitów
Pełna lista komend:
_joke, żart, hehe_`

var help = abstract.NewHelp(short, long)

func (j *joke) GetHelp() *abstract.Help {
	return help
}
