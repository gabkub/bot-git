package hello

import (
	"bot-git/bot/abstract"
	"bot-git/messageBuilders"
	"math/rand"
	"strings"
)

type hello struct {
}

var commands abstract.ReactForMsgs = []string{"cześć", "hej", "siema", "siemka", "siemanko", "hejo", "hejka", "elo"}

func New() *hello {
	return &hello{}
}

func (h *hello) CanHandle(msg string) bool {
	return commands.ContainsMessage(msg)
}

func (h *hello) Handle(msg string, sender abstract.MessageSender) {
	helloMsg := commands[rand.Intn(len(commands)-1)]
	text := strings.ToTitle(string(helloMsg[0])) + helloMsg[1:]
	sender.Send(messageBuilders.Text(text))
}
