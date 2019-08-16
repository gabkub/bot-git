package joke

import (
	"bot-git/bot/abstract"
	"bot-git/bot/commands/suchar"
	"bot-git/bot/jokes"
	"bot-git/bot/limit"
	"bot-git/messageBuilders"
	"bot-git/notNowMsg"
)

type joke struct {
}

var commands abstract.ReactForMsgs = []string{"joke", "żart", "hehe"}

func New() *joke {
	return &joke{}
}

func (j *joke) CanHandle(msg string) bool {
	return commands.ContainsMessage(msg)
}

func (j *joke) Handle(msg string, sender abstract.MessageSender) {
	if limit.CanGetJoke(sender.GetUserId()) {
		joke := jokes.Fetch(sender.GetUserId(), false)
		sentPost := sender.Send(messageBuilders.Text(joke))
		if sentPost != nil {
			suchar.SetLast(sentPost.Id)
		}
		return
	}
	sender.Send(messageBuilders.Text(notNowMsg.Get()))
}
