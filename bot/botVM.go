package bot

import (
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-server/model"
	"strings"
)
var bcfg config.MMConfig

func Start(ws *model.WebSocketClient, botConfig *config.BotConfig, mmCfg *config.MMConfig){


	//mod := &model.WebSocketEvent{}
	//Sets data
	Initialize()

	go func() {
		for {
			select {
			case resp := <-ws.EventChannel:
				HandleEvent(resp, mmCfg, botConfig)
			}
		}
	}()
	// You can block forever with
	select {}
}

func HandleEvent(event *model.WebSocketEvent, bc *config.MMConfig, cfg *config.BotConfig) {
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
		SendMsg(response, post.ChannelId)
	}
}


func SendMsg(msg, chId string) {
	post := &model.Post{}
	post.ChannelId = chId
	post.Message = msg

	if _, resp := bcfg.Client.CreatePost(post); resp.Error != nil {
		// log
		println("We failed to send a message to the logging channel")
		//PrintError(resp.Error)
	}
}


