package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"strings"
	"sync"
)

type alive struct {
	commands []string
	sync.Mutex
}

var A alive

func (a *alive) New() abstract.Handler {
	a.commands = []string{"alive","up","running", "żyjesz"}
	return a
}

func (a *alive) CanHandle(msg string) bool {
	return abstract.FindCommand(a.commands, msg)
}

func (a *alive) Handle(msg string) config.Msg {
	a.Lock()
	defer a.Unlock()

	if strings.Contains(msg, "-h") {
		return a.GetHelp()
	}
	toSend := config.Msg{"",config.Image{"Żyję!","https://media.giphy.com/media/6lK3ocoEWLFOo/giphy.gif"}, false}
	return toSend
}

func (a *alive) GetHelp() config.Msg {
	var sb strings.Builder
	sb.WriteString("Informacja, czy bot jest włączony i działa poprawnie.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_alive, up, running, żyjesz_\n")
	return config.Msg{sb.String(),config.Image{}, false}
}