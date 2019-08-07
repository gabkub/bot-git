package alive

import (
	"bot-git/bot/abstract"
	"bot-git/messageBuilders"
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

func (a *alive) Handle(msg string, sender abstract.MessageSender) {
	if strings.Contains(msg, "-h") {
		sender.Send(messageBuilders.Text(a.GetHelp()))
		return
	}
	sender.Send(messageBuilders.Text("Żyję!"))
}

func (a *alive) GetHelp() string {
	var sb strings.Builder
	sb.WriteString("Informacja, czy bot jest włączony i działa poprawnie.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_alive, up, running, żyjesz_\n")
	return sb.String()
}
