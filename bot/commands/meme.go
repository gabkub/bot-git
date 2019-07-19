package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/memes"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"strings"
	"sync"
)

type meme struct {
	commands []string
	sync.Mutex
}

var M meme

func (m *meme) New() abstract.Handler {
	m.commands = []string{"meme", "mem"}
	return m
}

func (m *meme) CanHandle(msg string) bool {
	return abstract.FindCommand(m.commands, msg)
}

func (m *meme) Handle(msg string) config.Msg {
	m.Lock()
	defer m.Unlock()

	if strings.Contains(msg, "-h") {
		return m.GetHelp()
	}
	if limit.CanSend(abstract.GetUserId(),"meme") {
		meme := memes.Fetch()
		return config.Msg{"", meme, false}
	}
	return abstract.LimitMsg()
}

func (m *meme) GetHelp() config.Msg {
	var sb strings.Builder
	sb.WriteString("Wysyła losowy śmieszny obrazek.\n")
	sb.WriteString("Atrybut -r usuwa ostatni mem.\n\n")
	sb.WriteString("Limity:\n")
	sb.WriteString("7:00-8:59 - 3 memy\n")
	sb.WriteString("9:00-14:59 - 1 mem na godzinę\n")
	sb.WriteString("15:00-6:59 - brak limitów\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_meme, mem_\n")
	return config.Msg{sb.String(),config.Image{}, false}
}