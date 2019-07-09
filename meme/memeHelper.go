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
		h.mainContainerID = rand.Intn(10)

		if h.mainContainerID != 0 && h.mainContainerID != lastID{
			h.mainContainer = fmt.Sprintf(h.mainContainer, h.mainContainerID)
			break
		}
	}

}
