package messages

import "github.com/mattermost/mattermost-server/model"

type Image struct{
	Header 		string
	ImageUrl 	string
}

func (i Image) IsEmpty() bool{
	return i.Header == "" && i.ImageUrl == ""
}

type Message struct {
	Text           string
	Img            Image
	Buttons        []*model.PostAction
	IsFunnyMessage bool
}

func (msg *Message) New() {
	msg.Text = ""
	msg.Img = Image{}
	msg.Buttons = nil
	msg.IsFunnyMessage = false
}

func (msg *Message) GetType() string {
	if !msg.Img.IsEmpty() {
		return "Image"
	}
	if msg.Buttons != nil {
		return "Buttons"
	}
	return "Text"
}
var Response Message
