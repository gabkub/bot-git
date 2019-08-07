package meme

import (
	"bot-git/bot/abstract"
	"bot-git/bot/limit"
	"bot-git/bot/memes"
	"bot-git/bot/messages"
	"bot-git/notNowMsg"
	"strings"
)

type meme struct {
	commands abstract.ReactForMsgs
}

func New() *meme {
	return &meme{[]string{"meme", "mem"}}
}

func (m *meme) CanHandle(msg string) bool {
	return m.commands.ContainsMessage(msg)
}

func (m *meme) Handle(msg string) messages.Message {
	if strings.Contains(msg, "-h") {
		return m.GetHelp()
	}
	if limit.CanSend(abstract.GetUserId(), "meme") {
		messages.Response.IsFunnyMessage = true
		meme := memes.Fetch()
		messages.Response.Img = meme
		return messages.Response
	}
	return notNowMsg.Get()
}

func (m *meme) GetHelp() messages.Message {
	var sb strings.Builder
	sb.WriteString("Wysyła losowy śmieszny obrazek. Odnośnik w tytule otwiera obrazek w nowej karcie.\n\n")
	sb.WriteString("Limity:\n")
	sb.WriteString("7:00-8:59 - 3 memy\n")
	sb.WriteString("9:00-14:59 - 1 mem na godzinę\n")
	sb.WriteString("15:00-6:59 - brak limitów\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_meme, mem_\n")
	messages.Response.Text = sb.String()
	return messages.Response
}
