package bot

import (
	"bot-git/bot/abstract"
	"bot-git/bot/commands/alive"
	"bot-git/bot/commands/football"
	"bot-git/bot/commands/hardJoke"
	"bot-git/bot/commands/hello"
	"bot-git/bot/commands/help"
	"bot-git/bot/commands/joke"
	"bot-git/bot/commands/meme"
	"bot-git/bot/commands/news"
	"bot-git/bot/commands/suchar"
	"bot-git/bot/commands/version"
	"bot-git/config"
	"bot-git/logg"
	"bot-git/messageBuilders"
	"bot-git/messageSender"
	"fmt"
	"github.com/mattermost/mattermost-server/model"
	"math/rand"
	"strings"
)

func handleEvent(event *model.WebSocketEvent) {
	// array of data from the event (user's message)
	post := model.PostFromJson(strings.NewReader(event.Data["post"].(string)))
	abstract.SetUserId(post.UserId)
	abstract.MsgChannel, _ = config.ConnectionCfg.Client.GetChannel(post.ChannelId, "")

	// ignore messages that are:
	// - empty
	// - bot's
	// - not sent directly to the bot

	prefix := fmt.Sprintf("@%s", config.BotCfg.BotName)
	if !canRespond(post, prefix) {
		return
	}
	m := strings.TrimSpace(strings.TrimPrefix(post.Message, prefix))
	sender := messageSender.New(post.ChannelId, getChannelType(post.ChannelId))
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

var defaultCommand = joke.New()

var handlers = []abstract.Handler{alive.New(), hello.New(), help.New(), defaultCommand,
	version.New(), meme.New(), suchar.New(), football.New(), news.New(),
	hardJoke.New()}

var gifs = []string{
	"https://media.giphy.com/media/pcOHEAG38BUaY/giphy.gif",
	"https://media.giphy.com/media/g7shkYchjuRBm/giphy.gif",
	"https://media.giphy.com/media/uL0pJDdA6fQ08/giphy.gif",
	"https://media.giphy.com/media/xzoXvpBoYTSKY/giphy.gif",
}

func handleMsg(msg string, sender abstract.MessageSender) {
	if msg == "" {
		img := gifs[rand.Intn(len(gifs))]
		sender.Send(messageBuilders.Image("Hello", img))
		return
	}
	for _, handler := range handlers {
		if handler.CanHandle(msg) {
			handler.Handle(msg, sender)
			return
		}
	}
	defaultCommand.Handle(msg, sender)
}
