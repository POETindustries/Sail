package store

import (
	"database/sql"
	"fmt"
	"sail/conf"
	"sail/errors"
	"strconv"
	"time"

	// mysql database driver
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct{}

func (m *mysql) Copy() Driver {
	return &mysql{}
}

func (m *mysql) Data(query *Query) []interface{} {
	return append(query.attrVals, query.selectionVals...)
}

func (m *mysql) Init() (*sql.DB, error) {
	return sql.Open("mysql", m.credentials())
}

func (m *mysql) Setup(table string, data []*SetupData) {
	var datatype, stmt string
	for _, d := range data {
		if m.asInt(d.Value, &datatype) {
			if d.IsPrimary {
				stmt += fmt.Sprintf("%s serial,", d.Name)
			} else {
				stmt += fmt.Sprintf("%s %s not null default %d,", d.Name, datatype, d.Value)
			}
		} else if m.asReal(d.Value, &datatype) {
			stmt += fmt.Sprintf("%s %s not null default %f,", d.Name, datatype, d.Value)
		} else if m.asText(d.Value, d.Size, &datatype) {
			if d.Size == All {
				stmt += fmt.Sprintf("%s %s,", d.Name, datatype)
			} else {
				stmt += fmt.Sprintf("%s %s not null default '%s',", d.Name, datatype, d.Value)
			}
		} else if m.asBool(d.Value, &datatype) {
			stmt += fmt.Sprintf("%s %s not null default %t,", d.Name, datatype, d.Value)
		} else if m.asTime(d.Value, &datatype) {
			stmt += fmt.Sprintf("%s %s not null default %d,", d.Name, datatype, d.Value.(time.Time).Unix())
		} else {
			return // TODO 2017-02-09: log unrecognized type error
		}
	}
	q := fmt.Sprintf("create table if not exists %s(%s)", table, stmt[:len(stmt)-1])
	if _, err := DB().Exec(q); err != nil {
		errors.Log(err, conf.Instance().DevMode)
	}
}

func (m *mysql) credentials() string {
	return conf.Instance().DBUser + ":" +
		conf.Instance().DBPass + "@tcp(" +
		conf.Instance().DBHost + ":3306)/" +
		conf.Instance().DBName + "?" +
		"tls=false"
}

func (m *mysql) asInt(v interface{}, outType *string) bool {
	switch v.(type) {
	case int8:
		*outType = "tinyint"
	case uint8:
		*outType = "tinyint unsigned"
	case int16:
		*outType = "smallint"
	case uint16:
		*outType = "smallint unsigned"
	case int32, int:
		*outType = "int"
	case uint32, uint:
		*outType = "int unsigned"
	case int64:
		*outType = "bigint"
	case uint64:
		*outType = "bigint unsigned"
	default:
		return false
	}
	return true
}

func (m *mysql) asReal(v interface{}, outType *string) bool {
	switch v.(type) {
	case float32:
		*outType = "float"
	case float64:
		*outType = "double"
	default:
		return false
	}
	return true
}

func (m *mysql) asBool(v interface{}, outType *string) bool {
	switch v.(type) {
	case bool:
		*outType = "bool"
		return true
	}
	return false
}

func (m *mysql) asText(v interface{}, size Datasize, outType *string) bool {
	switch v.(type) {
	case string:
		if size == All {
			*outType = "text"
		} else {
			*outType = "varchar(" + strconv.Itoa(int(size)) + ")"
		}
	case []string:
		*outType = "enum"
	case []byte:
		*outType = "longblob"
	default:
		return false
	}
	return true
}

func (m *mysql) asTime(v interface{}, outType *string) bool {
	switch v.(type) {
	case time.Time:
		*outType = "bigint"
		return true
	}
	return false
}
