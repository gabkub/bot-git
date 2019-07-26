package limit

import (
	"github.com/mattermost/mattermost-bot-sample-golang/config"
	"github.com/mattermost/mattermost-bot-sample-golang/logs"
	"log"
	"time"
)

const maxRequestsMorning = 3

type Limitation struct {
	Count        int
	LimitReached bool
}

var Users map[string]map[string]*Limitation

func getTeamId() string {
	team, resp := config.ConnectionCfg.Client.GetTeamByName(config.BotCfg.TeamName, "")
	if resp.Error != nil {
		log.Fatal("Error while getting team's ID. Details: " + resp.Error.DetailedError)
	}
	return team.Id
}

func SetUsersList() {
	teamMembers, resp := config.ConnectionCfg.Client.GetTeamMembers(getTeamId(),0,150,"")
	if resp.Error != nil {
		logs.WriteToFile("Error while getting team members'. Details: " + resp.Error.DetailedError)
	}

	Users = make(map[string]map[string]*Limitation)

	for _,user := range teamMembers {
		Users[user.UserId] = map[string]*Limitation{
			"joke": {0,false},
			"meme": {0,false},
			"football": {0,false},
		}
 	}
}


func AddRequest(userId, command string) {
	limit := Users[userId][command]
	limit.Count++
	limit.LimitReached = mustBlock(*limit)
	Users[userId][command] = limit
}

func mustBlock(limit Limitation) bool {
	hour := time.Now().Hour()
	if hour >= 7 && hour < 9 {
		if limit.Count >= maxRequestsMorning {
			return true
		}
	}
	if hour >=9 && hour < 15 {
		return true
	}
	return false
}

func CanSend(userId, command string) bool {
	return !Users[userId][command].LimitReached
}

