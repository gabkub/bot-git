package bot

import (
	"bot-git/bot/abstract"
	"bot-git/bot/commands/alive"
	"bot-git/bot/commands/football"
	"bot-git/bot/commands/hardJoke"
	"bot-git/bot/commands/hello"
	"bot-git/bot/commands/help"
	"bot-git/bot/commands/joke"
	"bot-git/bot/commands/meme"
	"bot-git/bot/commands/news"
	"bot-git/bot/commands/suchar"
	"bot-git/bot/commands/version"
)

var handlers = []abstract.Handler{alive.New(), hello.New(), helpHandler, defaultCommand,
	version.New(), meme.New(), suchar.New(), football.New(), news.New(),
	hardJoke.New(),
}

var defaultCommand = joke.New()
var helpHandler = help.New()
