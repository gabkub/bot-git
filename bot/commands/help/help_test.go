package help_test

import (
	"bot-git/bot/commands/help"
	"bot-git/testUtils/mockSender"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelp(t *testing.T) {
	h := help.New()
	sender := mockSender.New()

	h.Handle("help", sender)

	assert.NotEqual(t, "", sender.LastSentMsg)
}

func TestPomocy(t *testing.T) {
	h := help.New()
	sender := mockSender.New()

	h.Handle("pomocy", sender)

	assert.NotEqual(t, "", sender.LastSentMsg)
}
