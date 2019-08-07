package hello

import "bot-git/bot/abstract"

var short = "_cześć, hej_ - przywitanie"
var long = `Przywitanie :)

Pełna lista komend:
_cześć, hej, siema, siemanko, hejo, hejka, elo_`

var help = abstract.NewHelp(short, long)

func (h *hello) GetHelp() *abstract.Help {
	return help
}
