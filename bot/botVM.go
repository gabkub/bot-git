package bot

import (
	"../config"
	"github.com/mattermost/mattermost-server/model"
	"strings"
)

func Start(ws *model.WebSocketClient){

	go func() {
		for {
			select {
			case ev := <-ws.EventChannel:
				HandleEvent(ev)
			}
		}
	}()
	// You can block forever with
	select {}
}

func HandleEvent(event *model.WebSocketEvent) {
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
		if post.UserId == config.MmCfg.BotUser.Id{
			return
		}

		response := HandleMsg(post.Message)
		SendMsg(response, post.ChannelId)
	}
}


func SendMsg(msg, chId string) {
	post := &model.Post{}
	post.ChannelId = chId
	post.Message = msg

	if _, ev := config.MmCfg.Client.CreatePost(post); ev.Error != nil {
		// log
		println("We failed to send a message to the logging channel")
		//PrintError(resp.Error)
	}
}


