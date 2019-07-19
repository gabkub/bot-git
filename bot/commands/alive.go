package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"strings"
)

type alive struct {
	commands []string
}

var A alive

func (a *alive) New() abstract.Handler {
	a.commands = []string{"alive","up","running", "żyjesz"}
	return a
}

func (a *alive) CanHandle(msg string) bool {
	return abstract.FindCommand(a.commands, msg)
}

func (a *alive) Handle(msg string) messages.Message {
	if strings.Contains(msg, "-h") {
		return a.GetHelp()
	}
	messages.Response.Text = "Żyję!"
	return messages.Response
}

func (a *alive) GetHelp() messages.Message {
	var sb strings.Builder
	sb.WriteString("Informacja, czy bot jest włączony i działa poprawnie.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_alive, up, running, żyjesz_\n")
	messages.Response.Text = sb.String()
	return messages.Response
}