package hello

import (
	"bot-git/bot/abstract"
	"bot-git/messageBuilders"
	"math/rand"
	"strings"
)

type hello struct {
	commands abstract.ReactForMsgs
}

func New() *hello {
	return &hello{[]string{"cześć", "hej", "siema", "siemka", "siemanko", "hejo", "hejka", "elo"}}
}

func (h *hello) CanHandle(msg string) bool {
	return h.commands.ContainsMessage(msg)
}

func (h *hello) Handle(msg string, sender abstract.MessageSender) {
	if strings.Contains(msg, "-h") {
		sender.Send(messageBuilders.Text(h.GetHelp()))
		return
	}
	helloMsg := h.commands[rand.Intn(len(h.commands)-1)]
	text := strings.ToTitle(string(helloMsg[0])) + helloMsg[1:]
	sender.Send(messageBuilders.Text(text))
}

func (h *hello) GetHelp() string {
	var sb strings.Builder
	sb.WriteString("Przywitanie :)\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_cześć, hej, siema, siemanko, hejo, hejka, elo_\n")
	return sb.String()
}
