package tests

import (
	"fmt"
	"bot-git/bot/commands"
	"bot-git/config"
	"testing"
)

var msgAlive = config.Msg{"",config.Image{"Żyję!","https://media.giphy.com/media/6lK3ocoEWLFOo/giphy.gif"}, false}


func TestAlive(t *testing.T) {
	msg := commands.AliveHandler.Handle("alive")
	if msg != msgAlive {
		t.Error(fmt.Sprintf("Wrong response to 'alive'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsFunnyMessage))
	}
}

func TestUp(t *testing.T) {
	msg := commands.AliveHandler.Handle("up")
	if msg != msgAlive {
		t.Error(fmt.Sprintf("Wrong response to 'up'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsFunnyMessage))
	}
}
func TestRunning(t *testing.T) {
	msg := commands.AliveHandler.Handle("running")
	if msg != msgAlive {
		t.Error(fmt.Sprintf("Wrong response to 'running'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsFunnyMessage))
	}
}
func TestZyjesz(t *testing.T) {
	msg := commands.AliveHandler.Handle("żyjesz")
	if msg != msgAlive {
		t.Error(fmt.Sprintf("Wrong response to 'żyjesz'. Got:\nText: %s\nImage header: %s\nImage URL: %s\nIsJoke: %v",
			msg.Text, msg.Img.Header, msg.Img.ImageUrl,msg.IsFunnyMessage))
	}
}


