package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"strings"
)

type suchar struct {
	commands []string
}

var SucharHandler suchar
var lastFunnyMessage string

func (s *suchar) New() abstract.Handler {
	s.commands = []string{"suchar", "usuń", "delete", "no", "nie", "..."}
	return s
}

func (s *suchar) CanHandle(msg string) bool {
	return abstract.FindCommand(s.commands, msg)
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