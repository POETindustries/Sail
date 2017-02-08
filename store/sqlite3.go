package store

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
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

func (s *sqlite3) Setup(data *SetupData) {
	var schema, datatype, arg string
	for attr, val := range data.Data {
		if s.asText(val, &datatype) || s.asBlob(val, &datatype) {
			arg = "'%s'"
		} else if s.asInt(val, &datatype) {
			arg = "%d"
			if reflect.TypeOf(val).Kind() == reflect.Bool {
				val = s.btoi(val.(bool))
			}
		} else if s.asReal(val, &datatype) {
			arg = "%f"
		} else {
			return // TODO 2017-02-08: log unrecognized type error
		}
		if attr == data.Primary {
			schema += fmt.Sprintf("%s %s primary key not null,", attr, datatype)
		} else {
			schema += fmt.Sprintf("%s %s not null default "+arg+",", attr, datatype, val)
		}
	}
	query := fmt.Sprintf("create table if not exists %s(%s)", data.Relation, schema[:len(schema)-1])
	if _, err := DB().Exec(query); err != nil {
		errors.Log(err, conf.Instance().DevMode)
	}
}

func (s *sqlite3) credentials() string {
	return conf.Instance().DBName + ".db"
}

func (s *sqlite3) asText(v interface{}, res *string) bool {
	switch v.(type) {
	case string, time.Time:
		*res = "text"
		return true
	}
	return false
}

func (s *sqlite3) asInt(v interface{}, res *string) bool {
	switch v.(type) {
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, bool:
		*res = "integer"
		return true
	}
	return false
}

func (s *sqlite3) asReal(v interface{}, res *string) bool {
	switch v.(type) {
	case float32, float64:
		*res = "real"
		return true
	}
	return false
}

func (s *sqlite3) asBlob(v interface{}, res *string) bool {
	switch v.(type) {
	case []byte:
		*res = "blob"
		return true
	}
	return false
}

func (s *sqlite3) btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
