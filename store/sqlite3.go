package store

import (
	"database/sql"
	"os"
	"sail/conf"

	// sqlite database driver
	_ "github.com/mattn/go-sqlite3"
)

type sqlite3 struct{}

func (s *sqlite3) Copy() Driver {
	return &sqlite3{}
}

func (s *sqlite3) Init() (*sql.DB, error) {
	loc := conf.Instance().Cwd + "db/"
	if _, err := os.Stat(loc); err != nil {
		os.MkdirAll(loc, 0700)
	}
	return sql.Open("sqlite3", loc+"sail.db")
}

func (s *sqlite3) Param() string {
	return "?"
}

func (s *sqlite3) credentials() string {
	// sqlite does not implement access credentials
	return ""
}
