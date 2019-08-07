package football

import "bot-git/bot/abstract"

var short = "_piłkarzyki, gramy_ - rezerwacja stołu do gry w piłkarzyki"
var long = `Rezerwacja stołu do gry w piłkarzyki na 20 minut. Domyślna godzina rezerwacji to godzina wysłania wiadomości.
Limit rezerwacji na użytkownika = 1
Szablon: _<komenda>_ (@_<godzinarezerwacji>_) (domyślnie ustawiana jest aktualna godzina)
_<komenda>_ -l - wyświetla wszystkie rezerwacje na dany dzień.

Pełna lista komend:
_football, game, gramy, piłkarzyki, play, soccer_
`

var help = abstract.NewHelp(short, long)

func (f *football) GetHelp() *abstract.Help {
	return help
}
