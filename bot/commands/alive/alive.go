package alive

import (
	"bot-git/bot/abstract"
	"bot-git/bot/messages"
	"strings"
)

type alive struct {
	commands abstract.ReactForMsgs
}

func New() *alive {
	return &alive{[]string{"alive", "up", "running", "żyjesz"}}
}

func (a *alive) CanHandle(msg string) bool {
	return a.commands.ContainsMessage(msg)
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
