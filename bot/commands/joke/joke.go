package joke

import (
	"bot-git/bot/abstract"
	"bot-git/bot/jokes"
	"bot-git/bot/limit"
	"bot-git/bot/messages"
	"bot-git/notNowMsg"
	"strings"
)

type joke struct {
	commands abstract.ReactForMsgs
}

func New() *joke {
	return &joke{[]string{"joke", "żart", "hehe"}}
}

func (j *joke) CanHandle(msg string) bool {
	return j.commands.ContainsMessage(msg)
}

func (j *joke) Handle(msg string) messages.Message {

	if strings.Contains(msg, "-h") {
		return j.GetHelp()
	}
	if limit.CanSend(abstract.GetUserId(), "joke") {
		messages.Response.IsFunnyMessage = true
		joke := jokes.Fetch(false)
		messages.Response.Text = joke
		return messages.Response
	}
	return notNowMsg.Get()
}

func (j *joke) GetHelp() messages.Message {
	var sb strings.Builder
	sb.WriteString("Wysyła losowy dowcip. W dzień określony w pliku konfiguracyjnym żarty są w języku angielskim.\n")
	sb.WriteString("Limity:\n")
	sb.WriteString("7:00-8:59 - 3 żarty\n")
	sb.WriteString("9:00-14:59 - 1 żart na godzinę\n")
	sb.WriteString("15:00-6:59 - brak limitów\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_joke, żart, hehe_\n")
	messages.Response.Text = sb.String()
	return messages.Response
}
