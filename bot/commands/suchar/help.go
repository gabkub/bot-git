package suchar

import (
	"bot-git/bot/abstract"
)

var short = "_suchar, nie, ..._ - usuwa ostatni dowcip/mem"
var long = `Usuwa ostatni dowcip lub mem.
Pełna lista komend:
_suchar, usuń, delete, no, nie, ..._
`

var help = abstract.NewHelp(short, long)

func (s *suchar) GetHelp() *abstract.Help {
	return help
}
