package messages

// todo move to Fetch image
type Image struct {
	Header   string
	ImageUrl string
}

func NewImage(header, imageUrl string) *Image {
	return &Image{Header: header, ImageUrl: imageUrl}
}

func (i *Image) IsEmpty() bool {
	return i.Header == "" && i.ImageUrl == ""
}
