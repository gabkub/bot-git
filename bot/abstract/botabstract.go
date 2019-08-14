package abstract

import (
	"github.com/mattermost/mattermost-server/model"
	"strings"
)

type ReactForMsgs []string

func (r ReactForMsgs) ContainsMessage(msg string) bool {
	for _, v := range r {
		if strings.Contains(msg, v) {
			return true
		}
	}
	return false
}

type Help struct {
	Short string
	Long  string
}

func NewHelp(short, long string) *Help {
	return &Help{Short: short, Long: long}
}

type Handler interface {
	CanHandle(msg string) bool
	Handle(msg string, sender MessageSender)
	GetHelp() *Help
}

type MessageSender interface {
	Send(*model.Post) *model.Post
	IsDirectSend() bool
	GetUserId() UserId
}

type Image struct {
	Header   string
	ImageUrl string
}

type UserId string

func NewImage(header, imageUrl string) *Image {
	return &Image{Header: header, ImageUrl: imageUrl}
}

func (i *Image) IsEmpty() bool {
	return i.Header == "" && i.ImageUrl == ""
}
