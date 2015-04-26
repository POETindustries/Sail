// Package dbase handles database connections and connection credentials.
// Credentials are currently stored in-code as constants, but this is
// only temporary.
package dbase

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "sl_user"
const dbpass = "13381in651337"

// Open establishes a connection to the database. It only asks for the
// database's name, all other credentials should be stored in code or
// config files
//
// If this fails, it is usually because of wrong credentials or because
// there is no mysql service running or the service does not listen on
// the defualt port (usually 3306).
func Open(dbname string) *sql.DB {
	dataSource := dbuser + ":" + dbpass + "@/" + dbname
	if db, err := sql.Open("mysql", dataSource); err == nil {
		return db
	} else {
		println(err.Error())
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
	} else {
		println(err.Error())
	}
	return nil
}
