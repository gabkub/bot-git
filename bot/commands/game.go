package commands

import (
	"../../config"
	"../abstract"
	"strings"
)

type game struct {
	commands []string
}

var G game

func (g *game) New() abstract.Handler {
	g.commands = []string{"gramy", "game", "piłkarzyki"}
	return g
}

func (g *game) CanHandle(msg string) bool {
	return abstract.FindCommand(g.commands, msg)
}

func (g *game) Handle(msg string) (config.Msg, error) {
	if strings.Contains(msg, "-h") {
		return g.GetHelp()
	}

	return config.Msg{}, nil
}

func (g *game) GetHelp() (config.Msg, error) {
	v, e := abstract.Help("../../bot/commands/game_help.txt")
	toSend := config.Msg{v,config.Image{},false}
	return toSend, e
}