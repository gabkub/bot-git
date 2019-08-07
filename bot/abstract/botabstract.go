package abstract

import (
	"bot-git/bot/messages"
	"github.com/mattermost/mattermost-server/model"
	"strings"
)

var MsgChannel *model.Channel

type ReactForMsgs []string

func (r ReactForMsgs) ContainsMessage(msg string) bool {
	for _, v := range r {
		if strings.Contains(msg, v) {
			return true
		}
	}
	return false
}

type Handler interface {
	CanHandle(msg string) bool
	Handle(msg string) messages.Message
	GetHelp() messages.Message
}

var userId string

func GetUserId() string {
	return userId
}

func SetUserId(id string) {
	userId = id
}
