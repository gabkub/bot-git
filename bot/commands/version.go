package commands

import (
	"../abstract"
	"../../config"
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

func (v *version) Handle(msg string) (config.Msg, error) {
	if strings.Contains(msg, "-h") {
		return v.GetHelp()
	}
	return config.Msg{VER, config.Image{}}, nil
}

func (v *version) GetHelp() (config.Msg, error) {
	value, e :=	abstract.Help("../../bot/commands/version_help.txt")
	toSend := config.Msg{value,config.Image{}}
	return toSend, e
}