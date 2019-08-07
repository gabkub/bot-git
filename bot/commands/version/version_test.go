package version_test

import (
	"bot-git/bot/commands/version"
	"bot-git/messageBuilders"
	"bot-git/testUtils/mockSender"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedMsg = messageBuilders.Text(version.VER)

func TestVersion(t *testing.T) {
	v := version.New()
	sender := mockSender.New()

	v.Handle("version", sender)

	assert.Equal(t, expectedMsg, sender.LastSentMsg)
}

func TestWersja(t *testing.T) {
	v := version.New()
	sender := mockSender.New()

	v.Handle("wersja", sender)

	assert.Equal(t, expectedMsg, sender.LastSentMsg)
}

func TestVer(t *testing.T) {
	v := version.New()
	sender := mockSender.New()

	v.Handle("ver", sender)

	assert.Equal(t, expectedMsg, sender.LastSentMsg)
}
