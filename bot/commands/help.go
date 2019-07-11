package commands

import (
	"../../config"
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

func (h *help) Handle(msg string) (config.Msg, error) {
	v, e :=	abstract.Help("../../help.txt")
	toSend := config.Msg{v,config.Image{}}
	return toSend, e
}

func (h *help) GetHelp() (config.Msg, error) {
	v, e :=	abstract.Help("../../bot/commands/help_help.txt")
	toSend := config.Msg{v,config.Image{}}
	return toSend, e
}



