package meme

import (
	"bot-git/bot/abstract"
	"bot-git/bot/limit"
	"bot-git/bot/memes"
	"bot-git/messageBuilders"
	"bot-git/notNowMsg"
)

type meme struct {
}

var commands abstract.ReactForMsgs = []string{"meme", "mem"}

func New() *meme {
	return &meme{}
}

func (m *meme) CanHandle(msg string) bool {
	return commands.ContainsMessage(msg)
}

func (m *meme) Handle(msg string, sender abstract.MessageSender) {
	if limit.CanSend(abstract.GetUserId(), "meme") {
		meme := memes.Fetch()
		sender.Send(messageBuilders.Image(meme.Header, meme.ImageUrl))
		return
	}
	sender.Send(messageBuilders.Text(notNowMsg.Get()))
}

func (m *meme) GetHelp() *abstract.Help {
	return help
}
