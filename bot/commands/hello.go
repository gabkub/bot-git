package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
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

func (h *hello) Handle(msg string) config.Msg {
	h.Lock()
	defer h.Unlock()
	if strings.Contains(msg, "-h") {
		return h.GetHelp()
	}
	r := h.commands[rand.Intn(len(h.commands)-1)]
	return config.Msg{strings.ToTitle(string(r[0])) + r[1:], config.Image{},false}
}

func (h *hello) GetHelp() config.Msg {
	var sb strings.Builder
	sb.WriteString("Przywitanie :)\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_cześć, hej, siema, siemanko, hejo, hejka, elo_\n")
	return config.Msg{sb.String(),config.Image{},false}
}
