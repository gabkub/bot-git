package messageSender

import (
	"bot-git/bot/abstract"
	"bot-git/config"
	"bot-git/logg"
	"github.com/mattermost/mattermost-server/model"
)

const directType = "D"

type sender struct {
	channelId   string
	channelType string
	userId      abstract.UserId
}

func (s *sender) GetUserId() abstract.UserId {
	return s.userId
}

func New(userId abstract.UserId, channelId, channelType string) *sender {
	return &sender{userId: userId, channelId: channelId, channelType: channelType}
}

func (s *sender) IsDirectSend() bool {
	return s.channelType == directType
}

func (s *sender) Send(toSend *model.Post) *model.Post {
	if toSend != nil {
		toSend.ChannelId = s.channelId

		sentPost, er := config.ConnectionCfg.Client.CreatePost(toSend)
		if er.Error != nil {
			logg.WriteToFile("We failed to send a message to the logging channel. Details: " + er.Error.DetailedError)
			return nil
		}
		return sentPost
	} else {
		logg.WriteToFile("Error creating the respond message.")
	}
	return nil
}
