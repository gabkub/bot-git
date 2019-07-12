package commands

import (
	"../../config"
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
	//	return config.Msg{v, meme.Image{}}, e
	//}
	v, e := joker.Fetch()
	return config.Msg{v, config.Image{}}, e
}

func (j *joke) GetHelp() (config.Msg, error) {
	var sb strings.Builder
	sb.WriteString("Wysyła losowy dowcip. W dzień określony w pliku konfiguracyjnym żarty są w języku angielskim.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_joke, suchar, żart, hehe_\n")
	toSend := config.Msg{sb.String(),config.Image{}}
	return toSend, nil
}