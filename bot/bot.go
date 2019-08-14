package bot

import (
	"bot-git/bot/abstract"
	"bot-git/config"
	"bot-git/logg"
	"bot-git/messageBuilders"
	"bot-git/messageSender"
	"fmt"
	"github.com/mattermost/mattermost-server/model"
	"math/rand"
	"strings"
)

var gifs = []string{
	"https://media.giphy.com/media/pcOHEAG38BUaY/giphy.gif",
	"https://media.giphy.com/media/g7shkYchjuRBm/giphy.gif",
	"https://media.giphy.com/media/uL0pJDdA6fQ08/giphy.gif",
	"https://media.giphy.com/media/xzoXvpBoYTSKY/giphy.gif",
}

func handleEvent(event *model.WebSocketEvent) {
	// array of data from the event (user's message)
	post := model.PostFromJson(strings.NewReader(event.Data["post"].(string)))

	// ignore messages that are:
	// - empty
	// - bot's
	// - not sent directly to the bot

	prefix := fmt.Sprintf("@%s", config.BotCfg.BotName)
	if !canRespond(post, prefix) {
		return
	}
	m := strings.TrimSpace(strings.TrimPrefix(post.Message, prefix))
	sender := messageSender.New(abstract.UserId(post.UserId), post.ChannelId, getChannelType(post.ChannelId))
	handleMsg(m, sender)
	logg.WriteToFile("Message sent.")
}

func getChannelType(channelId string) string {
	ch, r := config.ConnectionCfg.Client.GetChannel(channelId, "")
	if r.Error != nil {
		return ""
	}
	return ch.Type

}

func canRespond(post *model.Post, prefix string) bool {
	post.Message = strings.ToLower(post.Message)
	return post != nil && post.UserId != config.ConnectionCfg.BotUser.Id && strings.Contains(post.Message, prefix)
}

func handleMsg(msg string, sender abstract.MessageSender) {
	if msg == "" {
		img := gifs[rand.Intn(len(gifs))]
		sender.Send(messageBuilders.Image("Hello", img))
		return
	}
	for _, handler := range handlers {
		if handler.CanHandle(msg) {
			if isHelpMsg(msg) {
				handleHelp(handler, sender)
				return
			}
			handler.Handle(msg, sender)
			return
		}
	}
	defaultCommand.Handle(msg, sender)
}

func handleHelp(handler abstract.Handler, sender abstract.MessageSender) {
	text := handler.GetHelp().Long
	sender.Send(messageBuilders.Text(text))
}

func isHelpMsg(msg string) bool {
	return strings.Contains(msg, "-h")
}
