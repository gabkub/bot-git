package commands

import (
	"../../config"
	"../abstract"
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

func (a *alive) Handle(msg string) (config.Msg, error) {
	if strings.Contains(msg, "-h") {
		return a.GetHelp()
	}
	toSend := config.Msg{"",config.Image{"Żyję!","https://media.giphy.com/media/6lK3ocoEWLFOo/giphy.gif"}, false}
	return toSend, nil
}

func (a *alive) GetHelp() (config.Msg, error) {
	var sb strings.Builder
	sb.WriteString("Informacja, czy bot jest włączony i działa poprawnie.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_alive, up, running, żyjesz_\n")
	toSend := config.Msg{sb.String(),config.Image{}, false}
	return toSend, nil
}