package alive

import "bot-git/bot/abstract"

var short = "_żyjesz, alive_ - sprawdza czy bot działa"
var long = `Informacja, czy bot jest włączony i działa poprawnie.

Pełna lista komend:
_alive, up, running, żyjesz_`

var help = abstract.NewHelp(short, long)

func (a *alive) GetHelp() *abstract.Help {
	return help
}
