package alive

import (
	"bot-git/bot/abstract"
	"bot-git/messageBuilders"
)

type alive struct {
}

var commands abstract.ReactForMsgs = []string{"alive", "up", "running", "żyjesz"}

func New() *alive {
	return &alive{}
}

func (a *alive) CanHandle(msg string) bool {
	return commands.ContainsMessage(msg)
}

func (a *alive) Handle(msg string, sender abstract.MessageSender) {
	sender.Send(messageBuilders.Text("Żyję!"))
}
