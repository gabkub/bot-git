package hardJoke

import (
	"bot-git/bot/abstract"
	"bot-git/bot/jokes"
	"bot-git/bot/limit"
	"bot-git/bot/messages"
	"bot-git/notNowMsg"
	"strings"
)

type hardJoke struct {
	commands abstract.ReactForMsgs
}

func New() abstract.Handler {
	return &hardJoke{[]string{"hard"}}
}

func (hj *hardJoke) CanHandle(msg string) bool {
	return hj.commands.ContainsMessage(msg)
}

func (hj *hardJoke) Handle(msg string) messages.Message {
	if strings.Contains(msg, "-h") {
		return hj.GetHelp()
	}
	if limit.CanSend(abstract.GetUserId(), "joke") {
		if abstract.MsgChannel.Type == "D" {
			messages.Response.IsFunnyMessage = true
			joke := jokes.Fetch(true)
			messages.Response.Text = joke
		} else {
			messages.Response.Text = "Tylko na priv."
		}
		return messages.Response
	}
	return notNowMsg.Get()
}

func (hj *hardJoke) GetHelp() messages.Message {
	var sb strings.Builder
	sb.WriteString("Wysyła losowy dowcip. Możliwe wylosowanie z kategorii **hard**. Komenda dostępna tylko w wiadomościach prywatnych z botem.\n")
	sb.WriteString("Żart **hard** nalicza się do ogólnego limitu żartów.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_hard_\n")
	messages.Response.Text = sb.String()
	return messages.Response
}
