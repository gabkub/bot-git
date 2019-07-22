package bot

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/logs"
	"github.com/mattermost/mattermost-server/model"
	"sync"
)

func Start(websocket *model.WebSocketClient){
	logs.WriteToFile("Bot has started.")

	go func() {
		for {
			select {
			case <- websocket.PingTimeoutChannel:
				logs.WriteToFile("Ping timeout.")
			case event := <-websocket.EventChannel:
				mux := &sync.Mutex{}
				mux.Lock()
				if websocket.ListenError != nil {
					logs.WriteToFile(fmt.Sprintf("ListenError occurred.\nWhere: %s\nStatus code: %v", websocket.ListenError.Where, websocket.ListenError.StatusCode))
				}
				if event.IsValid() && isMessage(event.Event) {
					handleEvent(event)
				}
				mux.Unlock()
			}
		}
	}()
	// block to the go function
	select {}
}

func isMessage(eventType string) bool {
	if eventType == model.WEBSOCKET_EVENT_POSTED {
		return true
	}
	return false
}

func sendMessage(channelId string, msg messages.Message) {
	// create new post
	var toSend *model.Post
	switch msg.GetType() {
	case "Text":
		toSend = &model.Post{
			Message: msg.Text,
		}
	case "Image":
		toSend = &model.Post{
			Message: msg.Text,
			Props: map[string]interface{}{
				"attachments": []model.SlackAttachment{
					{
						ImageURL: msg.Img.ImageUrl,
						Title: msg.Img.Header,},
				},
			},
		}
	}
	if toSend != nil {
		toSend.ChannelId = channelId

		sentPost, er := config.MmCfg.Client.CreatePost(toSend)
		if er.Error != nil {
			logs.WriteToFile("We failed to send a message to the logging channel. Details: " + er.Error.DetailedError)
		}

		if msg.IsJoke {
			commands.SetLastJoke(sentPost.Id)
		}
	} else {
		logs.WriteToFile("Error creating the respond message.")
	}
}