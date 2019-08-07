package version

import (
	"bot-git/bot/abstract"
	"bot-git/bot/messages"
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

func (v *version) Handle(msg string) messages.Message {
	if strings.Contains(msg, "-h") {
		return v.GetHelp()
	}
	messages.Response.Text = VER
	return messages.Response
}

func (v *version) GetHelp() messages.Message {
	var sb strings.Builder
	sb.WriteString("Zwraca aktualną wersję bota.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_wersja, version, ver_\n")
	messages.Response.Text = sb.String()
	return messages.Response
}
