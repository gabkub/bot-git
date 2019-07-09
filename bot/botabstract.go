package bot

import (
	"bufio"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/joker"
	"math/rand"
	"os"
	"strings"
)

type handler interface {
	CanHandle(msg string) bool
	Handle(cfg *config.BotConfig) (string, error)
}

type joke struct {
	commands []string
}

func (j joke) CanHandle(msg string) bool {
	return findCommand(j.commands, msg)
}

func (j joke) Handle(cfg *config.BotConfig) (string, error ) {
	return joker.Fetch(cfg)
}

type alive struct {
	commands []string
}

func (a alive) CanHandle(msg string) bool {
	return findCommand(a.commands, msg)
}

func (a alive) Handle(cfg *config.BotConfig) (string, error) {
	return "Żyję <3", nil
}

type help struct {
	commands []string
}

func (h help) CanHandle(msg string) bool {
	 return findCommand(h.commands, msg)
}

func (h help) Handle(cfg *config.BotConfig) (string, error) {
	return getHelp(cfg)
}

func getHelp(cfg *config.BotConfig) (string, error) {
	file, e := os.Open("help.txt")

	if e == nil {

		builder := strings.Builder{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			builder.WriteString(scanner.Text() + "\n")
		}

		return builder.String(), nil
	}
	return "Brak pliku pomocy", e
}

type hello struct {
	commands []string
}

func (h hello) CanHandle(msg string) bool {
	return findCommand(h.commands, msg)
}

func (h hello) Handle(cfg *config.BotConfig) (string, error) {
	r := h.commands[rand.Intn(len(h.commands)-1)]
	return strings.ToTitle(string(r[0])) + r[1:], nil
}

type version struct {
	commands []string
}

func (v version) CanHandle(msg string) bool {
	return findCommand(v.commands, msg)
}

func (v version) Handle(cfg *config.BotConfig) (string, error) {
	return VER, nil
}