package tests

import (
	"bot-git/bot/commands"
	"fmt"
	"testing"
)

func TestHelp(t *testing.T) {
	msg := commands.HelpHandler.Handle("help")
	if msg.Text == "" {
		t.Error(fmt.Sprintf("Wrong response to 'help'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl, msg.IsFunnyMessage))
	}
}

func TestPomocy(t *testing.T) {
	msg := commands.HelpHandler.Handle("pomocy")
	if msg.Text == "" {
		t.Error(fmt.Sprintf("Wrong response to 'pomocy'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl, msg.IsFunnyMessage))
	}
}
