package mockSender

import (
	"bot-git/bot/abstract"
	"github.com/mattermost/mattermost-server/model"
)

type MockSender struct {
	IsDirectedSendResult bool
	LastSentMsg          *model.Post
	SendResult           *model.Post
}

func (m *MockSender) GetUserId() abstract.UserId {
	return "1"
}

func New() *MockSender {
	return &MockSender{SendResult: &model.Post{}}
}

func (m *MockSender) Send(msg *model.Post) *model.Post {
	m.LastSentMsg = msg
	return m.SendResult
}

func (m *MockSender) IsDirectSend() bool {
	return m.IsDirectedSendResult
}
