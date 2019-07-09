package commands

import (
	"../abstract"
	"strings"
)

const VER = "1.0"

type version struct {
	commands []string
}

var V version

func (v *version) New() abstract.Handler {
	v.commands = []string{"wersja", "version", "ver"}
	return v
}

func (v *version) CanHandle(msg string) bool {
	return abstract.FindCommand(v.commands, msg)
}

func (v *version) Handle(msg string) (string, error) {
	if strings.Contains(msg, "-h") {
		return v.GetHelp()
	}
	return VER, nil
}

func (v *version) GetHelp() (string, error) {
	return abstract.Help("version_help.txt")
}