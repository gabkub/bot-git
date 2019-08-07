package joke

import (
	"bot-git/bot/abstract"
	"bot-git/bot/commands/suchar"
	"bot-git/bot/jokes"
	"bot-git/bot/limit"
	"bot-git/messageBuilders"
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

func (j *joke) Handle(msg string, sender abstract.MessageSender) {
	if strings.Contains(msg, "-h") {
		sender.Send(messageBuilders.Text(j.GetHelp()))
		return
	}
	if limit.CanSend(abstract.GetUserId(), "joke") {
		joke := jokes.Fetch(false)
		sentPost := sender.Send(messageBuilders.Text(joke))
		if sentPost != nil {
			suchar.SetLast(sentPost.Id)
		}
		return
	}
	sender.Send(messageBuilders.Text(notNowMsg.Get()))
}

func (j *joke) GetHelp() string {
	var sb strings.Builder
	sb.WriteString("Wysyła losowy dowcip. W dzień określony w pliku konfiguracyjnym żarty są w języku angielskim.\n")
	sb.WriteString("Limity:\n")
	sb.WriteString("7:00-8:59 - 3 żarty\n")
	sb.WriteString("9:00-14:59 - 1 żart na godzinę\n")
	sb.WriteString("15:00-6:59 - brak limitów\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_joke, żart, hehe_\n")
	return sb.String()
}
