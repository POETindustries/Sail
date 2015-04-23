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

func Open(dbname string) *sql.DB {
	dataSource := dbuser + ":" + dbpass + "@/" + dbname
	if db, err := sql.Open("mysql", dataSource); err == nil {
		return db
	} else {
		println(err.Error())
	}
	return nil
}
