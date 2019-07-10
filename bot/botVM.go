package bot

import (
	"../config"
	"fmt"
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
	prefix := fmt.Sprintf("@%s", config.BotCfg.Name)

	if !CanRespond(event, post, prefix) {
		return
	}

	response := HandleMsg(strings.TrimSpace(strings.TrimPrefix(post.Message, prefix)))
	SendMsg(post.ChannelId, response)
}

func CanRespond(event *model.WebSocketEvent, post *model.Post, prefix string) bool {
	post.Message = strings.ToLower(post.Message)
	if post == nil || post.UserId == config.MmCfg.BotUser.Id || !strings.Contains(post.Message, prefix) {
		return false
	}
	return true
}

func SendMsg(chId string, toSend config.Msg) {
	post := &model.Post{}

	//if !toSend.Img.IsEmpty() {
	//	post = &model.Post{
	//		Props: map[string]interface{}{
	//			"attachments": []model.SlackAttachment{
	//				{
	//					Color: "#7800FF",
	//					ImageURL: toSend.Img.ImageUrl,
	//					Title: toSend.Img.Header,
	//				},
	//			},
	//		},
	//	}}

	post.ChannelId = chId
	post.Message = toSend.Text

	if _, ev := config.MmCfg.Client.CreatePost(post); ev.Error != nil {
		// log
		println("We failed to send a message to the logging channel")
		//PrintError(resp.Error)
	}
}


