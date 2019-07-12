package commands

import (
	"../../config"
	"../abstract"
	"strings"
)

type help struct {
	commands []string
}

var H help

func (h *help) New() abstract.Handler {
	h.commands = []string{"help", "pomocy", "pomoc"}
	return h
}

func (h *help) CanHandle(msg string) bool {
	return abstract.FindCommand(h.commands, msg)
}

func (h *help) Handle(msg string) (config.Msg, error) {
	var sb strings.Builder
	sb.WriteString("LISTA KOMEND:\n")
	sb.WriteString(":arrow_right: _joke, żart_ - losowy dowcip\n")
	sb.WriteString(":arrow_right: _help, pomocy_ - pomoc\n")
	sb.WriteString(":arrow_right: _ver_ - wersja\n")
	sb.WriteString("<komenda> -h_ zwraca szczegółowe informacje o komendzie\n")
	toSend := config.Msg{sb.String(),config.Image{}}
	return toSend, nil
}

func (h *help) GetHelp() (config.Msg, error) {
	var sb strings.Builder
	sb.WriteString("Wyświetlenie ogólnej pomocy dla podstawowych komend\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_help, pomoc, pomocy_\n")
	toSend := config.Msg{sb.String(),config.Image{}}
	return toSend, nil
}



