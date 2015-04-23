package dbase

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const dbn = "pi_main"
const dbu = "pi_user"
const dbh = "127.0.0.1"
const dbp = "13381in651337"

func Open(name string) *sql.DB {
	dataSource := dbu + ":" + dbp + "@" + dbh + "/" + dbn
	if db, err := sql.Open("mysql", dataSource); err == nil {
		return db
	}
	return nil
}
