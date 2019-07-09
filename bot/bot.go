package bot

import (
	"./abstract"
	"./commands"
	"../config"
	"fmt"
	"strings"
)

// returns a response to the user if the command is one of the predefined commands

const DONTKNOW = "Nie rozumiem :( \nWpisz help, aby uzyskać listę komend."

func HandleMsg(msg string) string{
	prefix := fmt.Sprintf("@%s", config.BotCfg.Name)
	if strings.Contains(msg, prefix) {
		msg = strings.ToLower(strings.TrimSpace(strings.TrimPrefix(msg, prefix)))
		handlers := []abstract.Handler{commands.A.New(), commands.Hey.New(), commands.H.New(),  commands.J.New(), commands.V.New()}
		for _, hndl := range handlers {
			if hndl.CanHandle(msg) {
				if v, err := hndl.Handle(); err == nil {
					return v
				}
				return DONTKNOW
			}
		}
	}
	return DONTKNOW
}


