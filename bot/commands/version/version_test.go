package version_test

import (
	"bot-git/bot/commands/version"
	"bot-git/bot/messages"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedMsg = messages.Message{Text: version.VER, Img: messages.Image{}, IsFunnyMessage: false}

func TestVersion(t *testing.T) {
	v := version.New()

	msg := v.Handle("version")

	assert.Equal(t, expectedMsg, msg)
}

func TestWersja(t *testing.T) {
	v := version.New()

	msg := v.Handle("wersja")

	assert.Equal(t, expectedMsg, msg)
}

func TestVer(t *testing.T) {
	v := version.New()

	msg := v.Handle("ver")

	assert.Equal(t, expectedMsg, msg)
}
