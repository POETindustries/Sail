package store

import (
	"database/sql"
	"sail/conf"

	// mysql database driver
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct{}

func (m *mysql) Copy() Driver {
	return &mysql{}
}

func (m *mysql) Init() (*sql.DB, error) {
	return sql.Open("mysql", m.credentials())
}

func (m *mysql) Param() string {
	return "?"
}

func (m *mysql) credentials() string {
	return conf.Instance().DBUser + ":" +
		conf.Instance().DBPass + "@tcp(" +
		conf.Instance().DBHost + ":3306)/" +
		conf.Instance().DBName + "?" +
		"tls=false"
}