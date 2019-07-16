package bot

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-server/model"
	"log"
	"os"
)

func Start(websocket *model.WebSocketClient){

	limit.Start()
	log.Println("Bot has started.")

	go func() {
		for {
			select {
			case ev := <-websocket.EventChannel:
				if ev != nil {
				handleEvent(ev)}
			}
		}
	}()
	// block to the go function
	select {}
}

func sendMessage(channelId string, toSend config.Msg) {
	// create new post
	post := &model.Post{}

	// add attachments if needed
	if toSend.Img.ToAttach() {
		post = &model.Post{
			Props: map[string]interface{}{
				"attachments": []model.SlackAttachment{
					{
						ImageURL: toSend.Img.ImageUrl,
						Title: toSend.Img.Header,
					},
				},
			},
		}}

	post.ChannelId = channelId
	post.Message = toSend.Text
	p, er := config.MmCfg.Client.CreatePost(post)
	if er.Error != nil {
		log.Fatal("We failed to send a message to the logging channel. Details: " + er.Error.DetailedError)
		os.Exit(1)
	}
	// helper function for the removing of last joke functionality
	if toSend.IsJoke {
		commands.SetLastJoke(p.Id)
	}
}