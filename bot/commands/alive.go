package commands

import (
	"../abstract"
	"strings"
)

type alive struct {
	commands []string
}

var A alive

func (a *alive) New() abstract.Handler {
	a.commands = []string{"alive","up","running"}
	return a
}

func (a *alive) CanHandle(msg string) bool {
	return abstract.FindCommand(a.commands, msg)
}

func (a *alive) Handle(msg string) (string, error) {
	if strings.Contains(msg, "-h") {
		return a.GetHelp()
	}
	return "Żyję <3", nil
}

func (a *alive) GetHelp() (string, error) {
	return abstract.Help("alive_help.txt")
}