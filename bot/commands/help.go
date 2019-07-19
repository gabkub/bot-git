package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"strings"
	"sync"
)

type help struct {
	commands []string
	sync.Mutex
}

var H help

func (h *help) New() abstract.Handler {
	h.commands = []string{"help", "pomocy", "pomoc", "-h"}
	return h
}

func (h *help) CanHandle(msg string) bool {
	return abstract.FindCommand(h.commands, msg)
}

func (h *help) Handle(msg string) messages.Message {
	h.Lock()
	defer h.Unlock()
	var sb strings.Builder
	sb.WriteString("LISTA KOMEND:\n")
	sb.WriteString(":arrow_right: _joke, żart_ - losowy dowcip\n")
	sb.WriteString(":arrow_right: _meme, mem_ - losowy mem\n")
	sb.WriteString(":arrow_right: _help, pomocy_ - pomoc\n")
	sb.WriteString(":arrow_right: _ver_ - wersja\n")
	sb.WriteString("_<komenda> -h_ zwraca szczegółowe informacje o komendzie\n")
	messages.Response.Text = sb.String()
	return messages.Response
}

func (h *help) GetHelp() messages.Message {
	var sb strings.Builder
	sb.WriteString("Wyświetlenie ogólnej pomocy dla podstawowych komend\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_help, pomoc, pomocy_\n")
	messages.Response.Text = sb.String()
	return messages.Response
}



