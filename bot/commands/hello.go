package commands

import (
	"../../config"
	"../abstract"
	"math/rand"
	"strings"
)

type hello struct {
	commands []string
}

var Hey hello

func (h *hello) New() abstract.Handler {
	h.commands = []string{"cześć", "hej", "siema", "siemka", "siemanko", "hejo", "hejka", "elo"}
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
	return config.Msg{strings.ToTitle(string(r[0])) + r[1:], config.Image{},false}, nil
}

func (h *hello) GetHelp() (config.Msg, error) {
	var sb strings.Builder
	sb.WriteString("Przywitanie :)\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_cześć, hej, siema, siemanko, hejo, hejka, elo_\n")
	toSend := config.Msg{sb.String(),config.Image{},false}
	return toSend, nil
}
