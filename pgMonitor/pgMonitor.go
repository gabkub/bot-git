package pgMonitor

import (
	"bot-git/bot/messages"
	"bot-git/config"
	"bot-git/logg"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/mattermost/mattermost-server/model"
	"time"
)

type connection struct {
	pid          int
	database     sql.NullString
	userName     sql.NullString
	appName      sql.NullString
	clientAddr   sql.NullString
	backendStart pq.NullTime
	queryStart   pq.NullTime
	stateChange  pq.NullTime
	state        sql.NullString
	query        sql.NullString
}

type Connections []*connection

var alert messages.Message

func CheckCommand(event *model.WebSocketEvent) {

}

func CheckConnections() {
	alert.New()
	cons, err := getConnections()
	if err != nil {
		alert.Text = "Nie udało się pobrać połączeń z bazą. " + err.Error()
		//abstract.SendMessage(config.DbCfg.Channel.Id, alert)
		return
	}
	if cons != nil && len(cons) >= config.DbCfg.ConnectionsWarning {
		warning := fmt.Sprintf("Uwaga! Wysoka ilość połączeń z bazą %s (%d połączeń). Możesz je wylistować za pomocą komendy `zombie`", config.DbCfg.Name, len(cons))
		alert.Text = warning
		logg.WriteToFile(warning)
		logg.WriteToFile(formatConnections(cons))
	}
}

func LogConnections() {
	cons, err := getConnections()
	if err != nil {
		alert.Text = "Nie udało się pobrać połączeń z bazą. " + err.Error()
		//abstract.SendMessage(config.DbCfg.Channel.Id, alert)
		return
	}
	if cons != nil {
		header := fmt.Sprintf("\n\n%s Połączenia do %s:\n", time.Now().Format("15:04:05"), config.DbCfg.Name)
		logg.WriteToFile(header)
		logg.WriteToFile(formatConnections(cons))
	}
}
