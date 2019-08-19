package pgMonitor

import (
	"bot-git/config"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	getConnectionsSql = `SELECT pid, datname, usename, application_name, client_addr, backend_start, query_start, 
							state_change, state, query 
						FROM pg_stat_activity`
)

func getConnections() (Connections, error) {
	connStr := buildConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, newError("Błąd połączenia z %s : %s", config.DbCfg.Name, err.Error())
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	rows, err := db.Query(getConnectionsSql)
	if err != nil {
		return nil, newError("Błąd połączenia z %s : %s", config.DbCfg.Name, err.Error())
	}
	cons, err := scanConnections(rows)
	return cons, err
}

func buildConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DbCfg.Server, config.DbCfg.Port, config.DbCfg.User, config.DbCfg.Password, "template1")
}

func scanConnections(rows *sql.Rows) (Connections, error) {
	cons := make(Connections, 0, 0)
	for rows.Next() {
		con := &connection{}
		err := rows.Scan(&con.pid, &con.database, &con.userName, &con.appName, &con.clientAddr, &con.backendStart,
			&con.queryStart, &con.stateChange, &con.state, &con.query)
		if err != nil {
			return nil, newError("Błąd odczytu danych połączeń : %s", err.Error())
		}
		cons = append(cons, con)
	}
	return cons, nil
}

func newError(format string, values ...interface{}) error {
	return errors.New(fmt.Sprintf(format, values...))
}
