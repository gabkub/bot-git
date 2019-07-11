package commands

import (
	"../../config"
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

func (a *alive) Handle(msg string) (config.Msg, error) {
	if strings.Contains(msg, "-h") {
		return a.GetHelp()
	}
	toSend := config.Msg{"Żyję <3",config.Image{}}
	return toSend, nil
}

func (a *alive) GetHelp() (config.Msg, error) {
	v, e :=	abstract.Help("../../bot/commands/alive_help.txt")
	toSend := config.Msg{v,config.Image{}}
	return toSend, e
}