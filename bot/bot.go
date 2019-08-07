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
	"bot-git/bot/messages"
	"bot-git/config"
	"bot-git/logg"
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

	response := handleMsg(strings.TrimSpace(strings.TrimPrefix(post.Message, prefix)))
	SendMessage(post.ChannelId, response)
	logg.WriteToFile("Message sent.")
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

func handleMsg(msg string) messages.Message {
	messages.Response.New()
	if msg == "" {

		messages.Response.Img = messages.Image{Header: "Hello", ImageUrl: gifs[rand.Intn(len(gifs))]}
		return messages.Response
	}
	for _, handler := range handlers {
		if handler.CanHandle(msg) {
			return handler.Handle(msg)
		}

	}
	return defaultCommand.Handle(msg)
}
