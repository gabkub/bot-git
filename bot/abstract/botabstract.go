package abstract

import (
	"bufio"
	"../../config"
	"os"
	"strings"
)


type Handler interface {
	CanHandle(msg string) bool
	Handle(msg string) (config.Msg, error)
	GetHelp() (config.Msg, error)
}

func FindCommand(commands []string, msg string) bool {
	for _,v := range commands {
		if strings.Contains(msg, v){
			return true
		}
	}
	return false
}

func Help(path string) (string, error) {
	file, e := os.Open(path)

	if e == nil {
		builder := strings.Builder{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			builder.WriteString(scanner.Text() + "\n")
		}
		return builder.String(), nil
	}
	return "", e
}