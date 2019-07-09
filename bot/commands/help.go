package commands

import (
	"bufio"
	"../abstract"
	"os"
	"strings"
)

type help struct {
	commands []string
}

var H help

func (h *help) New() abstract.Handler {
	h.commands = []string{"help", "pomocy"}
	return h
}

func (h *help) CanHandle(msg string) bool {
	return abstract.FindCommand(h.commands, msg)
}

func (h *help) Handle() (string, error) {
	return getHelp()
}

func getHelp() (string, error) {
	file, e := os.Open("../help.txt")

	if e == nil {

		builder := strings.Builder{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			builder.WriteString(scanner.Text() + "\n")
		}

		return builder.String(), nil
	}
	return "Brak pliku pomocy", e
}

