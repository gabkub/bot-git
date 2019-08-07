package mockSender

import (
	"github.com/mattermost/mattermost-server/model"
)

type MockSender struct {
	IsDirectedSendResult bool
	LastSentMsg          *model.Post
	SendResult           *model.Post
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
