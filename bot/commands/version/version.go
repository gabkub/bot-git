package version

import (
	"bot-git/bot/abstract"
	"bot-git/messageBuilders"
)

const VER = "1.4.5"

type version struct {
}

var commands abstract.ReactForMsgs = []string{"wersja", "version", "ver"}

func New() *version {
	return &version{}
}

func (v *version) CanHandle(msg string) bool {
	return commands.ContainsMessage(msg)
}

func (v *version) Handle(msg string, sender abstract.MessageSender) {
	sender.Send(messageBuilders.Text(VER))
}
