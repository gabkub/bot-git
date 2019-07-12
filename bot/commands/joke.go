package commands

import (
	"../../config"
	"../abstract"
	"../jokes"
	"strings"
)

type joke struct {
	commands []string
}

var J joke
var LastJoke string

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
	if strings.Contains(msg, "-r") {
		return j.RemoveLast(), nil
	}
	//if abstract.FindCommand(j.commands[:3], msg) {
	//	v, e := jokes.Fetch()
	//	return config.Msg{v, meme.Image{}}, e
	//}
	v, e := jokes.Fetch()
	return config.Msg{v, config.Image{}, true}, e
}

func (j *joke) GetHelp() (config.Msg, error) {
	var sb strings.Builder
	sb.WriteString("Wysyła losowy dowcip. W dzień określony w pliku konfiguracyjnym żarty są w języku angielskim.\n")
	sb.WriteString("Atrybut -r usuwa ostatni żart.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_joke, suchar, żart, hehe_\n")
	toSend := config.Msg{sb.String(),config.Image{}, true}
	return toSend, nil
}

func (j *joke) RemoveLast() (config.Msg){
	if LastJoke == "" {
		return config.Msg{"Nie ma żartów do usunięcia...", config.Image{},false}
	}
	config.MmCfg.Client.DeletePost(LastJoke)
	return config.Msg{"", config.Image{"Done", "https://media.giphy.com/media/11lwLvxnaWobcc/giphy.gif"},false}
}