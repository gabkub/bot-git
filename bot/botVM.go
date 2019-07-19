package bot

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-server/model"
	"log"
	"sync"
)

func Start(websocket *model.WebSocketClient){
	log.Println("Bot has started.")

	go func() {
		for {
			select {
			case <- websocket.PingTimeoutChannel:
				log.Fatal("Ping timeout.")
			case ev := <-websocket.EventChannel:
				mux := &sync.Mutex{}
				mux.Lock()
				if websocket.ListenError != nil {
					log.Fatal("ListenError occurred. Reconnecting to websocket.")
				}
				if ev != nil {
				handleEvent(ev)}
				mux.Unlock()
			}
		}
	}()
	// block to the go function
	select {}
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
			log.Fatal("We failed to send a message to the logging channel. Details: " + er.Error.DetailedError)
		}

		if msg.IsJoke {
			commands.SetLastJoke(sentPost.Id)
		}
	} else {
		log.Fatal("Error creating the respond message.")
	}
}