package pgMonitor

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/ryanuber/columnize"
)

func formatConnections(cons Connections) string {
	if len(cons) == 0 {
		return ""
	}
	fields := "Pid | Database | UserName | AppName | ClientAddr | BackendStart | QueryStart | StateChange | State | Query"
	lines := make([]string, 0, 0)
	lines = append(lines, fields)
	for _, c := range cons {
		line := fmt.Sprintf("%d | %s | %s | %s | %s | %s | %s | %s | %s | %s",
			c.pid, c.database.String, c.userName.String, c.appName.String, c.clientAddr.String,
			formatTime(c.backendStart), formatTime(c.queryStart), formatTime(c.stateChange), c.state.String,
			limitString(c.query, 50))

		lines = append(lines, line)
	}
	return columnize.SimpleFormat(lines)
}

func formatTime(t pq.NullTime) string {
	return t.Time.Format("2006-01-02 15:04:05")
}

func limitString(s sql.NullString, limit int) string {
	if !s.Valid {
		return ""
	}
	if len(s.String) <= limit {
		return s.String
	}
	return s.String[:limit]
}