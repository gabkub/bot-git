package help_test

import (
	"bot-git/bot/commands/help"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelp(t *testing.T) {
	h := help.New()

	msg := h.Handle("help")

	assert.NotEqual(t, "", msg)
}

func TestPomocy(t *testing.T) {
	h := help.New()

	msg := h.Handle("pomocy")

	assert.NotEqual(t, "", msg)
}
