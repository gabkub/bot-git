package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/jokes"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"strings"
)

type hardJoke struct {
	commands []string
}

var HardJokeHandler hardJoke

func (hj *hardJoke) New() abstract.Handler {
	hj.commands = []string{"hard"}
	return hj
}

func (hj *hardJoke) CanHandle(msg string) bool {
	return abstract.FindCommand(hj.commands, msg)
}

func (hj *hardJoke) Handle(msg string) messages.Message {

	if strings.Contains(msg, "-h") {
		return hj.GetHelp()
	}
	if limit.CanSend(abstract.GetUserId(),"joke") {
		if abstract.MsgChannel.Type == "D" {
			messages.Response.IsFunnyMessage = true
			joke := jokes.Fetch(true)
			messages.Response.Text = joke
		} else {
			messages.Response.Text = "Tylko na priv."
		}
		return messages.Response
	}
	return abstract.RandomLimitMsg()
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