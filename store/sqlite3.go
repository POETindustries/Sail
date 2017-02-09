package store

import (
	"database/sql"
	"fmt"
	"os"
	"sail/conf"
	"sail/errors"
	"time"

	// sqlite database driver
	_ "github.com/mattn/go-sqlite3"
)

type sqlite3 struct{}

func (s *sqlite3) Copy() Driver {
	return &sqlite3{}
}

func (s *sqlite3) Data(query *Query) []interface{} {
	return append(query.attrVals, query.selectionVals...)
}

func (s *sqlite3) Init() (*sql.DB, error) {
	loc := conf.Instance().Cwd + "db/"
	if _, err := os.Stat(loc); err != nil {
		os.MkdirAll(loc, 0700)
	}
	return sql.Open("sqlite3", loc+s.credentials())
}

func (s *sqlite3) Setup(table string, data []*SetupData) {
	var stmt string
	for _, d := range data {
		switch d.Value.(type) {
		case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			if d.IsPrimary {
				stmt += fmt.Sprintf("%s integer primary key,", d.Name)
			} else {
				stmt += fmt.Sprintf("%s integer not null default %d,", d.Name, d.Value)
			}
		case float32, float64:
			stmt += fmt.Sprintf("%s real not null default %f,", d.Name, d.Value)
		case bool:
			stmt += fmt.Sprintf("%s integer not null default %d,", d.Name, s.btoi(d.Value.(bool)))
		case string:
			stmt += fmt.Sprintf("%s text not null default '%s',", d.Name, d.Value)
		case []byte:
			stmt += fmt.Sprintf("%s blob not null default '%s',", d.Name, d.Value)
		case time.Time:
			stmt += fmt.Sprintf("%s integer not null default %d,", d.Name, d.Value.(time.Time).Unix())
		default:
			return // TODO 2017-02-08: log unrecognized type error
		}
	}
	q := fmt.Sprintf("create table if not exists %s(%s)", table, stmt[:len(stmt)-1])
	if _, err := DB().Exec(q); err != nil {
		errors.Log(err, conf.Instance().DevMode)
	}
}

func (s *sqlite3) credentials() string {
	return conf.Instance().DBName + ".db"
}

func (s *sqlite3) btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
