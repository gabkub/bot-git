package bot

import (
	"../config"
	"./abstract"
	"./commands"
	"math/rand"
	"strings"
)

// returns a response to the user if the command is one of the predefined commands

func HandleMsg(msg string) config.Msg {
	handlers := []abstract.Handler{commands.A.New(), commands.Hey.New(), commands.H.New(),  commands.J.New(), commands.V.New()}
	for _, hndl := range handlers {
		if msg == "" {
			gifs := []string{
				"https://media.giphy.com/media/pcOHEAG38BUaY/giphy.gif",
				"https://media.giphy.com/media/g7shkYchjuRBm/giphy.gif",
				"https://media.giphy.com/media/uL0pJDdA6fQ08/giphy.gif",
				"https://media.giphy.com/media/xzoXvpBoYTSKY/giphy.gif",
			}
			return config.Msg{"", config.Image{"",gifs[rand.Intn(len(gifs))]},false}
		}
		if hndl.CanHandle(msg) {
			if strings.Contains(msg, "-h") {
				if v,e := hndl.GetHelp(); e == nil {
					return v
				}
			}
			if v, e := hndl.Handle(msg); e == nil {
				return v
			}
		}

	}
	def, _ := commands.J.Handle(msg)
	return def
}


