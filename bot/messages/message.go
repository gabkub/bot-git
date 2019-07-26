package messages

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
	Title		   string
	ThumbUrl		string
	IsFunnyMessage bool
}

func (msg *Message) New() {
	msg.Text = ""
	msg.Img = Image{}
	msg.Title = ""
	msg.ThumbUrl = ""
	msg.IsFunnyMessage = false
}

func (msg *Message) GetType() string {
	if !msg.Img.IsEmpty() {
		return "Image"
	}
	if msg.Title != "" && msg.ThumbUrl != "" {
		return "TitleThumbUrl"
	}
	if msg.Title != "" {
		return "Title"
	}
	if msg.ThumbUrl != "" {
		return "ThumbUrl"
	}
	return "Text"
}
var Response Message
