package bot

import (
	"fmt"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"strings"
)

// returns a response to the user if the command is one of the predefined commands
var j joke
var a alive
var h help
var powitanie hello
var ver version

const DONTKNOW = "Nie rozumiem :( \nWpisz help, aby uzyskać listę komend."
const VER = "1.0"

func Initialize(){

	j.commands = []string{"joke", "suchar", "żart", "hehe"}
	a.commands = []string{"alive","up","running"}
	h.commands = []string{"help", "pomocy"}
	powitanie.commands = []string{"cześć", "hej", "siema", "siemka", "siemanko", "hejo", "hejka"}
	ver.commands = []string{"wersja", "version", "ver"}
}

func HandleMsg(msg string, cfg *config.BotConfig) string{

	prefix := fmt.Sprintf("@%s", cfg.Name)
	if strings.Contains(msg, prefix) {

		msg = strings.TrimPrefix(msg, prefix)
		handlers := []handler{j, a, h, powitanie, ver}

		for _, hndl := range handlers {
			if hndl.CanHandle(msg) {

				if v, err := hndl.Handle(cfg); err == nil {
					return v
				}
				return DONTKNOW
			}
		}
	}
	return DONTKNOW
}

func findCommand(commands []string, msg string) bool {

	msg = strings.ToLower(strings.TrimSpace(msg))
	for _,v := range commands{
		if v == msg{
			return true
		}
	}
	return false
}

