package bot

import(
	"../joker"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
)
var aliveCommands = []string{"$alive","$up","$running"}
var jokeCommands = []string{"$joke", "$suchar", "$żart", "$hehe"}
// returns a response to the user if the command is one of the predefined commands
func HandleMsg(msg string, cfg *config.BotConfig) string{

	if FindCommand(aliveCommands, msg){
		return "<3 Tak, Żyję <3"
	} else if FindCommand(jokeCommands, msg){
		if joke, err := joker.Fetch(cfg); err ==nil {
			return joke
		}
	}
	return "Nie rozumiem :("

}

func FindCommand(commands []string, msg string) bool {

	for _,v := range commands{
		if v == msg{
			return true
		}
	}
	return false
}
