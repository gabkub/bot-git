package alive_test

import (
	"bot-git/bot/commands/alive"
	"bot-git/messageBuilders"
	"bot-git/testUtils/mockSender"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedMsg = messageBuilders.Text("Żyję!")

func TestAlive(t *testing.T) {
	a := alive.New()
	sender := mockSender.New()

	a.Handle("alive", sender)

	assert.Equal(t, expectedMsg, sender.LastSentMsg)
}

func TestUp(t *testing.T) {
	a := alive.New()
	sender := mockSender.New()

	a.Handle("up", sender)

	assert.Equal(t, expectedMsg, sender.LastSentMsg)
}
func TestRunning(t *testing.T) {
	a := alive.New()
	sender := mockSender.New()

	a.Handle("running", sender)

	assert.Equal(t, expectedMsg, sender.LastSentMsg)
}
func TestZyjesz(t *testing.T) {
	a := alive.New()
	sender := mockSender.New()

	a.Handle("żyjesz", sender)

	assert.Equal(t, expectedMsg, sender.LastSentMsg)
}
