package commands

import (
	"../abstract"
	"../../config"
	"strings"
)

const VER = "1.0.1.0"

type version struct {
	commands []string
}

var V version

func (v *version) New() abstract.Handler {
	v.commands = []string{"wersja", "version", "ver"}
	return v
}

func (v *version) CanHandle(msg string) bool {
	return abstract.FindCommand(v.commands, msg)
}

func (v *version) Handle(msg string) (config.Msg, error) {
	if strings.Contains(msg, "-h") {
		return v.GetHelp()
	}
	return config.Msg{VER, config.Image{},false}, nil
}

func (v *version) GetHelp() (config.Msg, error) {
	var sb strings.Builder
	sb.WriteString("Zwraca aktualną wersję bota.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_wersja, version, ver_\n")
	toSend := config.Msg{sb.String(),config.Image{},false}
	return toSend, nil
}