package bot

import "github.com/mattermost/mattermost-server/model"

type MMConfig struct{
	Client           *model.Client4
	WebSocketClient  *model.WebSocketClient
	BotUser          *model.User
	BotTeam          *model.Team
	BotConfig		 *model.Config
}