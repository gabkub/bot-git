package bot

import (
	"github.com/mattermost/mattermost-server/model"
	"strings"
	"../bot_sample.go"
)

var aliveCommands = [3]string{"$alive","$up","$running"}

// returns a response to the user if the command is one of the predefined commands
func Respond(msg  string) string{

	if FindCommand(aliveCommands, msg){
		return "<3 Tak, Żyję <3"
	} else {
		return "Nie rozumiem :("
	}

}

func FindCommand(commands [3]string, msg string) bool {

	for _,v := range commands{
		if v == msg{
			return true
		}
	}
	return false
}



func HandleMsg(event *model.WebSocketEvent) {

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
		if post.UserId == bot_sample.botUser.Id{
			return
		}
		CheckAlive(post.Message, post.Id, post.ChannelId)

	}
}

/*func SendMsg(msg string, replyToId, chId string) {
	post := &model.Post{}
	post.ChannelId = chId
	post.Message = msg

	post.RootId = replyToId

	if _, resp := client.CreatePost(post); resp.Error != nil {
		println("We failed to send a message to the logging channel")
		PrintError(resp.Error)
	}
}


func CheckAlive(msg , id, channelid string) {

	if FindCommand(aliveCommands, msg){
		SendMsg("<3 Tak żyję <3", id, channelid)
	} else {
		SendMsg("Nie rozumiem :(", id, channelid)
	}

	return
}


}*/

