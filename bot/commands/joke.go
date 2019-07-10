package commands

import (
	"../../config"
	"../../meme"
	"../../joker"
	"../abstract"
	"strings"
)

type joke struct {
	commands []string
}

var J joke

func (j *joke) New() abstract.Handler {
	j.commands = []string{"joke", "suchar", "żart", "hehe"} // TODO: dodać do helpa
	return j
}

func (j *joke) CanHandle(msg string) bool {
	return abstract.FindCommand(j.commands, msg)
}

func (j *joke) Handle(msg string) (config.Msg, error) {
	if strings.Contains(msg, "-h") {
		return j.GetHelp()
	}
	//if abstract.FindCommand(j.commands[:3], msg) {
	//	v, e := joker.Fetch()
	//	return config.Msg{v, meme.Meme{}}, e
	//}
	v, e := joker.Fetch()
	return config.Msg{v, meme.Meme{}}, e
}

func (j *joke) GetHelp() (config.Msg, error) {
	v, e := abstract.Help("../bot/commands/joke_help.txt")
	toSend := config.Msg{v,meme.Meme{}}
	return toSend, e
}