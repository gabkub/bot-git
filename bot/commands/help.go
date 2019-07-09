package commands

import (
	"../abstract"
)

type help struct {
	commands []string
}

var H help

func (h *help) New() abstract.Handler {
	h.commands = []string{"help", "pomocy", "pomoc"}
	return h
}

func (h *help) CanHandle(msg string) bool {
	return abstract.FindCommand(h.commands, msg)
}

func (h *help) Handle(msg string) (string, error) {
	return abstract.Help("../help.txt")
}

func (h *help) GetHelp() (string, error) {
	return abstract.Help("../bot/commands/help_help.txt")
}



