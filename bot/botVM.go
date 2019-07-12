package bot

import (
	"../config"
	"fmt"
	"./commands"
	"github.com/mattermost/mattermost-server/model"
	"os"
	"strings"
)

func Start(ws *model.WebSocketClient){
	println("Bot is now able to respond.")
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
	if !toSend.Img.IsEmpty() {
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

	post.ChannelId = chId
	post.Message = toSend.Text
	p, ev := config.MmCfg.Client.CreatePost(post)
	if ev.Error != nil {
		// log
		println("We failed to send a message to the logging channel")
		os.Exit(1)
		//PrintError(resp.Error)
	}
	if toSend.IsJoke {
		commands.LastJoke = p.Id
	}
}