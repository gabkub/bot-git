package bot

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/logs"
	"github.com/mattermost/mattermost-bot-sample-golang/main/connection"
	"github.com/mattermost/mattermost-bot-sample-golang/schedule"
	"github.com/mattermost/mattermost-server/model"
	"log"
	"sync"
)

var mux = &sync.Mutex{}

func Start(){

	logs.WriteToFile("Bot has started.")
	log.Println("Bot has started.")

	go func() {
		schedule.Start()
		for {
			select {

			case <-connection.Websocket.PingTimeoutChannel:
				mux.Lock()
				logs.WriteToFile("Websocket PingTimeout.")
				config.ConnectionCfg.Client.Logout()
				connection.Connect()
				mux.Unlock()

			case event := <-connection.Websocket.EventChannel:
				mux.Lock()
				if event != nil {
					if event.IsValid() && isMessage(event.Event){
						handleEvent(event)
					}
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
						Title: msg.Img.Header,
						TitleLink: msg.Img.ImageUrl,
					},

				},
			},
		}
	}
	if toSend != nil {
		toSend.ChannelId = channelId

		sentPost, er := config.ConnectionCfg.Client.CreatePost(toSend)
		if er.Error != nil {
			logs.WriteToFile("We failed to send a message to the logging channel. Details: " + er.Error.DetailedError)
		}

		if msg.IsFunnyMessage {
			commands.SetLast(sentPost.Id)
		}

	} else {
		logs.WriteToFile("Error creating the respond message.")
	}
}