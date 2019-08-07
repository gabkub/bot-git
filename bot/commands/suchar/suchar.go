package suchar

import (
	"bot-git/bot/abstract"
	"bot-git/bot/messages"
	"bot-git/config"
	"strings"
)

type suchar struct {
	commands abstract.ReactForMsgs
}

var lastFunnyMessage string

func New() *suchar {
	return &suchar{[]string{"suchar", "usuń", "delete", "no", "nie", "..."}}
}

func (s *suchar) CanHandle(msg string) bool {
	return s.commands.ContainsMessage(msg)
}

func (s *suchar) Handle(msg string) messages.Message {

	if strings.Contains(msg, "-h") {
		return s.GetHelp()
	}
	return s.removeLast()
}

func (s *suchar) GetHelp() messages.Message {
	var sb strings.Builder
	sb.WriteString("Usuwa ostatni dowcip lub mem.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_suchar, usuń, delete, no, nie, ..._\n")
	messages.Response.Text = sb.String()
	return messages.Response
}

func (s *suchar) removeLast() messages.Message {
	if lastFunnyMessage == "" {
		messages.Response.Text = "Nie wiem, o co ci chodzi..."
		return messages.Response
	}
	config.ConnectionCfg.Client.DeletePost(lastFunnyMessage)
	messages.Response.Text = "ok"
	return messages.Response
}

func SetLast(last string) {
	lastFunnyMessage = last
}
