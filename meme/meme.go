package meme

type Meme struct{
	Header 		string
	ImageUrl 	string
}

var mem meme

func Fetch() Meme {

	mem.getMemedroid("https://www.memedroid.com/memes/top/day")

	return Meme{mem.header, mem.imageUrl}
}

func (m Meme)IsEmpty() bool{

	if m.Header == "" && m.ImageUrl == ""{
		return true
	}
	return false
}