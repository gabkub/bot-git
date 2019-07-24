package pgMonitor

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/mattermost/mattermost-bot-sample-golang/config"
)

type connection struct {
	pid           int
	database      sql.NullString
	userName      sql.NullString
	appName       sql.NullString
	clientAddr    sql.NullString
	backendStart  pq.NullTime
	queryStart    pq.NullTime
	stateChange   pq.NullTime
	state         sql.NullString
	query         sql.NullString
}

type Connections []*connection

func checkDBConnections(dbConfig *config.DatabaseConfig) {
	//connections, err := getConnections(dbConfig)
}

func getConnections(dbConfig *config.DatabaseConfig) (Connections, error) {

}