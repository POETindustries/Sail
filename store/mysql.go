/******************************************************************************
Copyright 2015-2017 POET Industries

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the
following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
******************************************************************************/

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

func (m *mysql) Prepare(query string) string {
	return query
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
