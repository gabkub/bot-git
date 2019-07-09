package commands

import (
	"../../joker"
	"../abstract"
	"strings"
)

type joke struct {
	commands []string
}

var J joke

func (j *joke) New() abstract.Handler {
	j.commands = []string{"joke", "suchar", "żart", "hehe"}
	return j
}

func (j *joke) CanHandle(msg string) bool {
	return abstract.FindCommand(j.commands, msg)
}

func (j *joke) Handle(msg string) (string, error ) {
	if strings.Contains(msg, "-h") {
		return j.GetHelp()
	}
	return joker.Fetch()
}

func (j *joke) GetHelp() (string, error) {
	return abstract.Help("joke_help.txt")
}