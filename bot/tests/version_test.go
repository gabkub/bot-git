package tests

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"testing"
)

var msgVersion = config.Msg{commands.VER, config.Image{},false}

func TestVersion(t *testing.T) {
	msg := commands.V.Handle("version")
	if msg != msgVersion {
		t.Error(fmt.Sprintf("Wrong response to 'version'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsJoke))
	}
}

func TestWersja(t *testing.T) {
	msg := commands.V.Handle("wersja")
	if msg != msgVersion {
		t.Error(fmt.Sprintf("Wrong response to 'wersja'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsJoke))
	}
}

func TestVer(t *testing.T) {
	msg := commands.V.Handle("ver")
	if msg != msgVersion {
		t.Error(fmt.Sprintf("Wrong response to 'ver'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsJoke))
	}
}

