package limit

import (
"github.com/mattermost/mattermost-bot-sample-golang/config"
"log"
"os"
	"time"
)

const maxRequestsMorning = 3

type Limitation struct {
	Count        int
	LimitReached bool
}

var Users map[string]map[string]*Limitation

func getTeamId() string {
	team, resp := config.MmCfg.Client.GetTeamByName(config.BotCfg.TeamName, "")
	if resp.Error != nil {
		log.Fatal("Error while getting team's ID. Details: " + resp.Error.DetailedError)
		os.Exit(1)
	}
	return team.Id
}

func SetTeamMembers () {
	teamMembers, resp := config.MmCfg.Client.GetTeamMembers(getTeamId(),0,150,"")
	if resp.Error != nil {
		log.Fatal("Error while getting team members'. Details: " + resp.Error.DetailedError)
		os.Exit(1)
	}

	Users = make(map[string]map[string]*Limitation)

	for _,user := range teamMembers {
		Users[user.UserId] = map[string]*Limitation{
			"joke": {0,false},
			"meme": {0,false},
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