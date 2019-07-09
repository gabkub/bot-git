package commands

import (
	"../abstract"
	"../../config"
	"../../meme"
	"math/rand"
	"strings"
)

type hello struct {
	commands []string
}

var Hey hello

func (h *hello) New() abstract.Handler {
	h.commands = []string{"cześć", "hej", "siema", "siemka", "siemanko", "hejo", "hejka"}
	return h
}

func (h *hello) CanHandle(msg string) bool {
	return abstract.FindCommand(h.commands, msg)
}

func (h *hello) Handle(msg string) (config.Msg, error) {
	if strings.Contains(msg, "-h") {
		return h.GetHelp()
	}
	r := h.commands[rand.Intn(len(h.commands)-1)]
	return config.Msg{strings.ToTitle(string(r[0])) + r[1:], meme.Meme{}}, nil
}

func (h *hello) GetHelp() (config.Msg, error) {
	v, e :=	abstract.Help("../bot/commands/hello_help.txt")
	toSend := config.Msg{v,meme.Meme{}}
	return toSend, e
}
