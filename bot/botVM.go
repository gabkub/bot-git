package bot

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/footballDatabase"
	"github.com/mattermost/mattermost-bot-sample-golang/logg"
	"github.com/mattermost/mattermost-bot-sample-golang/main/connection"
	"github.com/mattermost/mattermost-bot-sample-golang/schedule"
	"github.com/mattermost/mattermost-server/model"
	"log"
	"sync"
)

var mux = &sync.Mutex{}

func Start(){

	logg.WriteToFile("Bot has started.")
	log.Println("Bot has started.")

	go func() {
		schedule.Start()
		footballDatabase.CreateTableDB()
		for {
			select {

			case <-connection.Websocket.PingTimeoutChannel:
				mux.Lock()
				logg.WriteToFile("Websocket PingTimeout.")
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

func SendMessage(channelId string, msg messages.Message) {
	// create new post
	var toSend *model.Post
	switch msg.GetType() {
	case "Text":
		toSend = &model.Post{
			Message: msg.Text,
		}
	case "News":
		toSend = &model.Post{
			Props: map[string]interface{}{
				"attachments": []model.SlackAttachment{
					{
						ImageURL: msg.Img.ImageUrl,
						Title: msg.Img.Header,
						TitleLink: msg.TitleLink,
					},
				},
			},
		}
	case "Image":
		toSend = &model.Post{
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

	case "Title":
		toSend = &model.Post{
			Message: msg.Text,
			Props: map[string]interface{}{
				"attachments": []model.SlackAttachment{
					{
						Title: msg.Title,
					},

				},
			},
		}

	case "ThumbUrl":
		toSend = &model.Post{
			Message: msg.Text,
			Props: map[string]interface{}{
				"attachments": []model.SlackAttachment{
					{
						ThumbURL: msg.ThumbUrl,
					},

				},
			},
		}

	case "TitleThumbUrl":
		toSend = &model.Post{
			Props: map[string]interface{}{
				"attachments": []model.SlackAttachment{
					{
						Fields: []*model.SlackAttachmentField{
								{
									Short: false,
									Title: msg.Title,
									Value: msg.Text,
								},
						},
						ThumbURL: msg.ThumbUrl,
					},

				},
			},
		}
	}

	if toSend != nil {
		toSend.ChannelId = channelId

		sentPost, er := config.ConnectionCfg.Client.CreatePost(toSend)
		if er.Error != nil {
			logg.WriteToFile("We failed to send a message to the logging channel. Details: " + er.Error.DetailedError)
		}

		if msg.IsFunnyMessage {
			commands.SetLast(sentPost.Id)
		}

	} else {
		logg.WriteToFile("Error creating the respond message.")
	}
}