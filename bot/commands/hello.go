package commands

import (
	"../abstract"
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

func (h *hello) Handle(msg string) (string, error) {
	if strings.Contains(msg, "-h") {
		return h.GetHelp()
	}
	r := h.commands[rand.Intn(len(h.commands)-1)]
	return strings.ToTitle(string(r[0])) + r[1:], nil
}

func (h *hello) GetHelp() (string, error) {
	return abstract.Help("hello_help.txt")
}
