package limit

import (
	"bot-git/bot/abstract"
	"bot-git/config"
	"bot-git/logg"
	"log"
	"time"
)

const (
	maxRequestsMorning   = 3
	maxRequestsDuringDay = 1
	meme                 = "meme"
	joke                 = "joke"
)

type limitation struct {
	count int
}

var users map[abstract.UserId]map[string]*limitation

func getTeamId() string {
	team, resp := config.ConnectionCfg.Client.GetTeamByName(config.BotCfg.TeamName, "")
	if resp.Error != nil {
		log.Fatal("Error while getting team's ID. Details: " + resp.Error.DetailedError)
	}
	return team.Id
}

func setUsersList() {
	teamMembers, resp := config.ConnectionCfg.Client.GetTeamMembers(getTeamId(), 0, 150, "")
	if resp.Error != nil {
		logg.WriteToFile("Error while getting team members'. Details: " + resp.Error.DetailedError)
	}

	users = make(map[abstract.UserId]map[string]*limitation)

	for _, user := range teamMembers {
		users[abstract.UserId(user.UserId)] = map[string]*limitation{
			joke: {0},
			meme: {0},
		}
	}
}

func AddJoke(userId abstract.UserId) {
	addRequest(userId, joke)
}

func AddMeme(userId abstract.UserId) {
	addRequest(userId, meme)
}

func addRequest(userId abstract.UserId, command string) {
	users[userId][command].count++
}

func (l *limitation) LimitReached() bool {
	hour := time.Now().Hour()
	if hour >= 7 && hour < 9 {
		if l.count >= maxRequestsMorning {
			return true
		}
	}
	if hour >= 9 && hour < 15 {
		return l.count >= maxRequestsDuringDay
	}
	return false
}

func CanGetMeme(userId abstract.UserId) bool {
	return canSend(userId, meme)
}

func CanGetJoke(userId abstract.UserId) bool {
	return canSend(userId, joke)
}

func canSend(userId abstract.UserId, command string) bool {
	return !users[userId][command].LimitReached()
}

func Reset() {
	for _, user := range users {
		for _, userLimit := range user {
			userLimit.count = 0
		}
	}
}

func InitUser() {
	if users == nil {
		setUsersList()
	}
}
