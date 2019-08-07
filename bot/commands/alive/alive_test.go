package alive_test

import (
	"bot-git/bot/commands/alive"
	"bot-git/bot/messages"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedMsg = messages.Message{Text: "Żyję!", Img: messages.Image{}, IsFunnyMessage: false}

func TestAlive(t *testing.T) {
	a := alive.New()

	msg := a.Handle("alive")

	assert.Equal(t, expectedMsg, msg)
}

func TestUp(t *testing.T) {
	a := alive.New()

	msg := a.Handle("up")

	assert.Equal(t, expectedMsg, msg)
}
func TestRunning(t *testing.T) {
	a := alive.New()

	msg := a.Handle("running")

	assert.Equal(t, expectedMsg, msg)
}
func TestZyjesz(t *testing.T) {
	a := alive.New()

	msg := a.Handle("żyjesz")

	assert.Equal(t, expectedMsg, msg)
}
