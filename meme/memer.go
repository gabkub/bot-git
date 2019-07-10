package meme

type Meme struct{
	Header 		string
	ImageUrl 	string
}

const maxRequest  = 3
var countRequest int
var mem meme
func Fetch() Meme  {

	if countRequest >= maxRequest{
		return Meme{"Weź się do roboty.", ""}
	}

	mem.getMemedroid("https://www.memedroid.com/memes/top/day")
	countRequest++

	return Meme{mem.header, mem.imageUrl}
}

func (m Meme)IsEmpty() bool{
	
	if m.Header == "" && m.ImageUrl == ""{
		return true
	}
	return false
}