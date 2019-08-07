package version

import (
	"bot-git/bot/abstract"
)

var short = "_ver_ - wersja"
var long = `Zwraca aktualną wersję bota.
Pełna lista komend:
_wersja, version, ver_
`

var help = abstract.NewHelp(short, long)

func (v *version) GetHelp() *abstract.Help {
	return help
}
