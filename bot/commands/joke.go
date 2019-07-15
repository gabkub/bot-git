package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/jokes"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"strings"
	"sync"
)

type joke struct {
	commands []string
	sync.Mutex
}

var J joke
var lastJoke string

func (j *joke) New() abstract.Handler {
	j.commands = []string{"joke", "suchar", "żart", "hehe"}
	return j
}

func (j *joke) CanHandle(msg string) bool {
	return abstract.FindCommand(j.commands, msg)
}

func (j *joke) Handle(msg string) config.Msg {
	j.Lock()
	defer j.Unlock()

	if strings.Contains(msg, "-h") {
		return j.GetHelp()
	}
	if strings.Contains(msg, "-r") {
		return j.removeLast()
	}
	joke := jokes.Fetch()
	return config.Msg{joke, config.Image{}, true}
}

func (j *joke) GetHelp() config.Msg {
	var sb strings.Builder
	sb.WriteString("Wysyła losowy dowcip. W dzień określony w pliku konfiguracyjnym żarty są w języku angielskim.\n")
	sb.WriteString("Atrybut -r usuwa ostatni żart.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_joke, suchar, żart, hehe_\n")
	return config.Msg{sb.String(),config.Image{}, true}
}

func (j *joke) removeLast() config.Msg {
	if lastJoke == "" {
		return config.Msg{"Nie ma żartów do usunięcia...", config.Image{},false}
	}
	config.MmCfg.Client.DeletePost(lastJoke)
	return config.Msg{"", config.Image{"Done", "https://media.giphy.com/media/11lwLvxnaWobcc/giphy.gif"},false}
}

func SetLastJoke(last string) {
	lastJoke = last
}