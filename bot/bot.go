package bot

import (
	"../config"
	"./abstract"
	"./commands"
	"strings"
)

// returns a response to the user if the command is one of the predefined commands

// const DONTKNOW = "Nie rozumiem :( \nWpisz help, aby uzyskać listę komend."

func HandleMsg(msg string) config.Msg {
	handlers := []abstract.Handler{commands.A.New(), commands.Hey.New(), commands.H.New(),  commands.J.New(), commands.V.New()}
	for _, hndl := range handlers {
		if hndl.CanHandle(msg) {
			if strings.Contains(msg, "-h") {
				if v,e := hndl.GetHelp(); e == nil {
					return v
				}
			}
			if v, e := hndl.Handle(msg); e == nil {
				return v
			}
			return config.Msg{"Błąd", config.Image{}}
		}
	}
	v, _ := commands.J.Handle(msg)
	return v
}


