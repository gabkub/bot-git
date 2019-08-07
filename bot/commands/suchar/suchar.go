package suchar

import (
	"bot-git/bot/abstract"
	"bot-git/config"
	"bot-git/messageBuilders"
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

func (s *suchar) Handle(msg string, sender abstract.MessageSender) {
	if strings.Contains(msg, "-h") {
		sender.Send(messageBuilders.Text(s.GetHelp()))
		return
	}
	text := s.removeLast()
	sender.Send(messageBuilders.Text(text))
}

func (s *suchar) GetHelp() string {
	var sb strings.Builder
	sb.WriteString("Usuwa ostatni dowcip lub mem.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_suchar, usuń, delete, no, nie, ..._\n")
	return sb.String()
}

func (s *suchar) removeLast() string {
	if lastFunnyMessage == "" {
		return "Nie wiem, o co ci chodzi..."
	}
	config.ConnectionCfg.Client.DeletePost(lastFunnyMessage)
	return "ok"
}

func SetLast(last string) {
	lastFunnyMessage = last
}
