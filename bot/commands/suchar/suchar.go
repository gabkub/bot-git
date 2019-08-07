package suchar

import (
	"bot-git/bot/abstract"
	"bot-git/config"
	"bot-git/messageBuilders"
)

type suchar struct {
}

var commands abstract.ReactForMsgs = []string{"suchar", "usuń", "delete", "no", "nie", "..."}
var lastFunnyMessage string

func New() *suchar {
	return &suchar{}
}

func (s *suchar) CanHandle(msg string) bool {
	return commands.ContainsMessage(msg)
}

func (s *suchar) Handle(msg string, sender abstract.MessageSender) {
	text := s.removeLast()
	sender.Send(messageBuilders.Text(text))
}

func (s *suchar) removeLast() string {
	if lastFunnyMessage == "" {
		return "Nie wiem, o co ci chodzi..."
	}
	config.ConnectionCfg.Client.DeletePost(lastFunnyMessage)
	return "ok"
}

func SetLast(last string) {
	lastFunnyMessage = last
}
