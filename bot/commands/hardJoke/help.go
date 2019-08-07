package hardJoke

import "bot-git/bot/abstract"

var short = "_hard_ (dostępne tylko w wiadomościach prywatnych z botem) - losowy żart w kategorii **hard** (na własną odpowiedzialność!)"
var long = `Wysyła losowy dowcip. Możliwe wylosowanie z kategorii **hard**. Komenda dostępna tylko w wiadomościach prywatnych z botem.
Żart **hard** nalicza się do ogólnego limitu żartów.

Pełna lista komend:
_hard_`

var help = abstract.NewHelp(short, long)

func (hj *hardJoke) GetHelp() *abstract.Help {
	return help
}
