package bot

import (
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-server/model"
	"strings"
)
var bcfg MMConfig
func Start(event *model.WebSocketEvent, bc *MMConfig, cfg *config.BotConfig) {
	bcfg = *bc
	// Lets only reponded to messaged posted events
	if event.Event != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	// log
	// println("responding to debugging channel msg")

	// array of data from the event (user's message)
	post := model.PostFromJson(strings.NewReader(event.Data["post"].(string)))
	if post != nil {

		// ignore my events
		if post.UserId == bc.BotUser.Id{
			return
		}

		response := HandleMsg(post.Message, cfg)
		SendMsg(response, post.Id, post.ChannelId)
	}
}

func SendMsg(msg string, replyToId, chId string) {
	post := &model.Post{}
	post.ChannelId = chId
	post.Message = msg

	post.RootId = replyToId

	if _, resp := bcfg.Client.CreatePost(post); resp.Error != nil {
		// log
		println("We failed to send a message to the logging channel")
		//PrintError(resp.Error)
	}
}


