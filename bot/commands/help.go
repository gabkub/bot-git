package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"strings"
)

type help struct {
	commands []string
}

var HelpHandler help

func (h *help) New() abstract.Handler {
	h.commands = []string{"help", "pomocy", "pomoc"}
	return h
}

func (h *help) CanHandle(msg string) bool {
	return abstract.FindCommand(h.commands, msg)
}

func (h *help) Handle(msg string) messages.Message {
	var sb strings.Builder
	sb.WriteString("LISTA KOMEND:\n")
	sb.WriteString("- _joke, żart_ - losowy dowcip\n")
	sb.WriteString("- _meme, mem_ - losowy mem\n")
	sb.WriteString("- _suchar, nie, ..._ - usuwa ostatni dowcip/mem\n")
	sb.WriteString("- _help, pomocy_ - pomoc\n")
	sb.WriteString("- _ver_ - wersja\n")
	sb.WriteString("- _<komenda> -h_ zwraca szczegółowe informacje o komendzie\n")
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



