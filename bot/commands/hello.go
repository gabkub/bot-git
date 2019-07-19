package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"math/rand"
	"strings"
	"sync"
)

type hello struct {
	commands []string
	sync.Mutex
}

var Hey hello

func (h *hello) New() abstract.Handler {
	h.commands = []string{"cześć", "hej", "siema", "siemka", "siemanko", "hejo", "hejka", "elo"}
	return h
}

func (h *hello) CanHandle(msg string) bool {
	return abstract.FindCommand(h.commands, msg)
}


func (h *hello) Handle(msg string) messages.Message {
	h.Lock()
	defer h.Unlock()

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
