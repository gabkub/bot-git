package meme

import "../config"

type Image config.Image

const maxRequest  = 3
var countRequest int
var mem meme
func Fetch() Image {

	if countRequest >= maxRequest{
		return Image{"Weź się do roboty.", ""}
	}

	mem.getMemedroid("https://www.memedroid.com/memes/top/day")
	countRequest++

	return Image{mem.header, mem.imageUrl}
}