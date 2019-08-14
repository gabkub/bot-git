package hardJoke

import (
	"bot-git/bot/abstract"
	"bot-git/bot/commands/suchar"
	"bot-git/bot/jokes"
	"bot-git/bot/limit"
	"bot-git/messageBuilders"
	"bot-git/notNowMsg"
)

type hardJoke struct {
}

var commands abstract.ReactForMsgs = []string{"hard"}

func New() abstract.Handler {
	return &hardJoke{}
}

func (hj *hardJoke) CanHandle(msg string) bool {
	return commands.ContainsMessage(msg)
}

func (hj *hardJoke) Handle(msg string, sender abstract.MessageSender) {
	text, ok := getMessage(sender.GetUserId(), sender.IsDirectSend())
	sentPost := sender.Send(messageBuilders.Text(text))
	if ok && sentPost != nil {
		suchar.SetLast(sentPost.Id)
	}
}

func getMessage(userId abstract.UserId, isDirect bool) (string, bool) {
	const ok = true
	if limit.CanSend(userId, "joke") {
		if isDirect {
			joke := jokes.Fetch(userId, true)
			return joke, ok
		} else {
			return "Tylko na priv.", !ok
		}
	}
	return notNowMsg.Get(), !ok
}
