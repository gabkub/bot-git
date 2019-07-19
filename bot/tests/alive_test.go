package tests

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/commands"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"testing"
)

var msgAlive = config.Msg{"",config.Image{"Żyję!","https://media.giphy.com/media/6lK3ocoEWLFOo/giphy.gif"}, false}


func TestAlive(t *testing.T) {
	msg := commands.A.Handle("alive")
	if msg != msgAlive {
		t.Error(fmt.Sprintf("Wrong response to 'alive'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsJoke))
	}
}

func TestUp(t *testing.T) {
	msg := commands.A.Handle("up")
	if msg != msgAlive {
		t.Error(fmt.Sprintf("Wrong response to 'up'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsJoke))
	}
}
func TestRunning(t *testing.T) {
	msg := commands.A.Handle("running")
	if msg != msgAlive {
		t.Error(fmt.Sprintf("Wrong response to 'running'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsJoke))
	}
}
func TestZyjesz(t *testing.T) {
	msg := commands.A.Handle("żyjesz")
	if msg != msgAlive {
		t.Error(fmt.Sprintf("Wrong response to 'żyjesz'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsJoke))
	}
}


