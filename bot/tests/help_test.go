package tests

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"testing"
)

func TestHelp(t *testing.T) {
	msg := commands.H.Handle("help")
	if msg.Text=="" {
		t.Error(fmt.Sprintf("Wrong response to 'help'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsJoke))
	}
}

func TestPomocy(t *testing.T) {
	msg := commands.H.Handle("pomocy")
	if msg.Text=="" {
		t.Error(fmt.Sprintf("Wrong response to 'pomocy'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsJoke))
	}
}
