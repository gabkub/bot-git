package version

import (
	"bot-git/bot/abstract"
	"bot-git/messageBuilders"
	"strings"
)

const VER = "1.4.5"

type version struct {
	commands abstract.ReactForMsgs
}

func New() *version {
	return &version{[]string{"wersja", "version", "ver"}}
}

func (v *version) CanHandle(msg string) bool {
	return v.commands.ContainsMessage(msg)
}

func (v *version) Handle(msg string, sender abstract.MessageSender) {
	if strings.Contains(msg, "-h") {
		sender.Send(messageBuilders.Text(v.GetHelp()))
		return
	}
	sender.Send(messageBuilders.Text(VER))
}

func (v *version) GetHelp() string {
	var sb strings.Builder
	sb.WriteString("Zwraca aktualną wersję bota.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_wersja, version, ver_\n")
	return sb.String()
}
