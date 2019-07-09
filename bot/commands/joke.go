package commands

import (
	"../abstract"
	"../../joker"
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

func (j *joke) Handle() (string, error ) {
	return joker.Fetch()
}
