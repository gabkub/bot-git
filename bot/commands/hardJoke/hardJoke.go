package hardJoke

import (
	"bot-git/bot/abstract"
	"bot-git/bot/commands/suchar"
	"bot-git/bot/jokes"
	"bot-git/bot/limit"
	"bot-git/messageBuilders"
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

func (hj *hardJoke) Handle(msg string, sender abstract.MessageSender) {
	if strings.Contains(msg, "-h") {
		sender.Send(messageBuilders.Text(hj.GetHelp()))
		return
	}
	text, ok := getMessage(sender.IsDirectSend())
	sentPost := sender.Send(messageBuilders.Text(text))
	if ok && sentPost != nil {
		suchar.SetLast(sentPost.Id)
	}
}

func getMessage(isDirect bool) (string, bool) {
	const ok = true
	if limit.CanSend(abstract.GetUserId(), "joke") {
		if isDirect {
			joke := jokes.Fetch(true)
			return joke, ok
		} else {
			return "Tylko na priv.", !ok
		}
	}
	return notNowMsg.Get(), !ok
}

func (hj *hardJoke) GetHelp() string {
	var sb strings.Builder
	sb.WriteString("Wysyła losowy dowcip. Możliwe wylosowanie z kategorii **hard**. Komenda dostępna tylko w wiadomościach prywatnych z botem.\n")
	sb.WriteString("Żart **hard** nalicza się do ogólnego limitu żartów.\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_hard_\n")
	return sb.String()
}
