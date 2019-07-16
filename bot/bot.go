package bot

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"fmt"
	"github.com/mattermost/mattermost-server/model"
	"math/rand"
	"strings"
)

// returns a response to the user if the command is one of the predefined commands

func handleMsg(msg string) config.Msg {
	// initialize the handlers
	handlers := []abstract.Handler{commands.A.New(), commands.Hey.New(), commands.H.New(),  commands.J.New(),
		commands.V.New(), commands.M.New()}
	for _, hndl := range handlers {
		if msg == "" {
			gifs := []string{
				"https://media.giphy.com/media/pcOHEAG38BUaY/giphy.gif",
				"https://media.giphy.com/media/g7shkYchjuRBm/giphy.gif",
				"https://media.giphy.com/media/uL0pJDdA6fQ08/giphy.gif",
				"https://media.giphy.com/media/xzoXvpBoYTSKY/giphy.gif",
			}
			return config.Msg{"", config.Image{"Hello",gifs[rand.Intn(len(gifs))]},false}
		}
		if hndl.CanHandle(msg) {
			return hndl.Handle(msg)
		}

	}
	return commands.J.Handle(msg)
}

func handleEvent(event *model.WebSocketEvent) {
	// ignore events which are not messages
	if event.Event != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	// array of data from the event (user's message)
	post := model.PostFromJson(strings.NewReader(event.Data["post"].(string)))
	abstract.SetUserId(post.UserId)
	// ignore messages that are:
	// - empty
	// - bot's
	// - not to the bot
	prefix := fmt.Sprintf("@%s", config.BotCfg.Name)
	if !canRespond(post, prefix) {
		return
	}
	// respond to the message
	response := handleMsg(strings.TrimSpace(strings.TrimPrefix(post.Message, prefix)))
	sendMessage(post.ChannelId, response)
}

func canRespond(post *model.Post, prefix string) bool {
	post.Message = strings.ToLower(post.Message)
	return post != nil && post.UserId != config.MmCfg.BotUser.Id && strings.Contains(post.Message, prefix)
}
