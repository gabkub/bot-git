package meme

import (
	"fmt"
	"math/rand"
	"time"
)

type helpFindMeme struct{
	mainContainer 		string
	mainContainerID 	int
	header        		string
	image         		string
}

func (h *helpFindMeme) randContainer () {
	rand.Seed(time.Now().UTC().UnixNano())

	for{
		lastID := h.mainContainerID
		h.mainContainerID = rand.Intn(30)
		println(lastID)
		println(h.mainContainerID)

		if h.mainContainerID != 0 && h.mainContainerID != lastID{
			h.mainContainer = fmt.Sprintf("article.gallery-item:nth-child(%d) div.item-aux-container", h.mainContainerID)
			break
		}
	}

}