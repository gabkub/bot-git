package commands

import (
	"bot-git/bot/abstract"
	"bot-git/bot/messages"
	"math/rand"
	"strings"
)

type hello struct {
	commands []string
}

var HelloHandler hello

func (h *hello) New() abstract.Handler {
	h.commands = []string{"cześć", "hej", "siema", "siemka", "siemanko", "hejo", "hejka", "elo"}
	return h
}

func (h *hello) CanHandle(msg string) bool {
	return abstract.FindCommand(h.commands, msg)
}


func (h *hello) Handle(msg string) messages.Message {

	if strings.Contains(msg, "-h") {
		return h.GetHelp()
	}
	helloMsg := h.commands[rand.Intn(len(h.commands)-1)]
	messages.Response.Text = strings.ToTitle(string(helloMsg[0])) + helloMsg[1:]
	return messages.Response
}

func (h *hello) GetHelp() messages.Message {
	var sb strings.Builder
	sb.WriteString("Przywitanie :)\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_cześć, hej, siema, siemanko, hejo, hejka, elo_\n")
	messages.Response.Text = sb.String()
	return messages.Response
}
