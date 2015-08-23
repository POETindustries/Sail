package dbase

import (
	"database/sql"
	"sail/conf"

	_ "github.com/lib/pq" // driver for postgres
)

// New creates a connection object to access the database.
//
// If this fails, it is usually because of wrong credentials or because
// there is no sql service running or the service does not listen on
// the default port.
func New(config *conf.Config) *Conn {
	conn := &Conn{
		Credentials: config.DBCredString(),
		DevMode:     config.DevMode}

	if conn.init() == nil && conn.Verify() {
		return conn
	}

	return nil
}

// QueryRow is a wrapper around the sql function of the same name. It takes
// a query string and the query parameters, constructs a statement out of
// them, checks for connection problems and returns one row and an error
// reference
func QueryRow(query string, db *sql.DB, args ...interface{}) *sql.Row {
	if stmt, err := db.Prepare(query); err == nil {
		return stmt.QueryRow(args...)
	}

	return nil
}
