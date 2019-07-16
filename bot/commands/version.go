package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"strings"
	"sync"
)

const VER = "1.0.2.2"

type version struct {
	commands []string
	sync.Mutex
}

var V version

func (v *version) New() abstract.Handler {
	v.commands = []string{"wersja", "version", "ver"}
	return v
}

func (v *version) CanHandle(msg string) bool {
	return abstract.FindCommand(v.commands, msg)
}

func (v *version) Handle(msg string) config.Msg {
	v.Lock()
	defer v.Unlock()

	if strings.Contains(msg, "-h") {
		return v.GetHelp()
	}
	return config.Msg{VER, config.Image{},false}
}

func (v *version) GetHelp() config.Msg {
	var sb strings.Builder
	sb.WriteString("Zwraca aktualną wersję bota.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_wersja, version, ver_\n")
	return config.Msg{sb.String(),config.Image{},false}
}