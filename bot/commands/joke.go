package commands

import (
	"github.com/mattermost/mattermost-bot-sample-golang/bot/abstract"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/jokes"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/limit"
	"github.com/mattermost/mattermost-bot-sample-golang/bot/messages"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"strings"
)

type joke struct {
	commands []string
}

var J joke
var lastJoke string

func (j *joke) New() abstract.Handler {
	j.commands = []string{"joke", "suchar", "żart", "hehe"}
	return j
}

func (j *joke) CanHandle(msg string) bool {
	return abstract.FindCommand(j.commands, msg)
}

func (j *joke) Handle(msg string) messages.Message {

	if strings.Contains(msg, "-h") {
		return j.GetHelp()
	}
	if strings.Compare(msg, "suchar") == 0 {
		return j.removeLast()
	}
	if limit.CanSend(abstract.GetUserId(),"joke") {
		messages.Response.IsJoke = true
		joke := jokes.Fetch()
		messages.Response.Text = joke
		return messages.Response
	}
	return abstract.LimitMsg()
}

func (j *joke) GetHelp() messages.Message {
	var sb strings.Builder
	sb.WriteString("Wysyła losowy dowcip. W dzień określony w pliku konfiguracyjnym żarty są w języku angielskim.\n")
	sb.WriteString("Komenda _suchar_ usuwa ostatni żart.\n\n")
	sb.WriteString("Limity:\n")
	sb.WriteString("7:00-8:59 - 3 żarty\n")
	sb.WriteString("9:00-14:59 - 1 żart na godzinę\n")
	sb.WriteString("15:00-6:59 - brak limitów\n\n")
	sb.WriteString("Pełna lista komend:\n")
	sb.WriteString("_joke, żart, hehe_\n")
	messages.Response.Text = sb.String()
	return messages.Response
}

func (j *joke) removeLast() messages.Message {
	if lastJoke == "" {
		messages.Response.Text = "Brak żartów do usunięcia..."
		return messages.Response
	}
	config.MmCfg.Client.DeletePost(lastJoke)
	messages.Response.Img.ImageUrl = "https://media.giphy.com/media/11lwLvxnaWobcc/giphy.gif"
	return messages.Response
}

func SetLastJoke(last string) {
	lastJoke = last
}