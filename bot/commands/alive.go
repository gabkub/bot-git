package commands

import (
	"../abstract"
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

func (a *alive) Handle() (string, error) {
	return "Żyję <3", nil
}
