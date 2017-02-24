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
	"sail/log"
	"strconv"
	"strings"
	"time"

	// postgres database driver
	_ "github.com/lib/pq"
)

type postgres struct{}

func (p *postgres) Copy() Driver {
	return &postgres{}
}

func (p *postgres) Data(query *Query) []interface{} {
	count := 0
	for i, a := range query.attrs {
		if strings.HasSuffix(a, "?") {
			count++
			query.attrs[i] = strings.Replace(a, "?", "$"+strconv.Itoa(count), 1)
		}
	}
	for i, s := range query.selection {
		if strings.HasSuffix(s, "?") {
			count++
			query.selection[i] = strings.Replace(s, "?", "$"+strconv.Itoa(count), 1)
		}
	}
	return append(query.attrVals, query.selectionVals...)
}

func (p *postgres) Init() (*sql.DB, error) {
	return sql.Open("postgres", p.credentials())
}

func (p *postgres) Prepare(query string) string {
	count := 0
	for strings.Contains(query, "?") {
		count++
		query = strings.Replace(query, "?", "$"+strconv.Itoa(count), 1)
	}
	return query
}

func (p *postgres) Setup(table string, data []*SetupData) {
	var datatype, stmt string
	for _, d := range data {
		if p.asInt(d.Value, &datatype) {
			if d.IsPrimary {
				stmt += fmt.Sprintf("%s bigserial,", d.Name)
			} else {
				stmt += fmt.Sprintf("%s %s not null default %d,", d.Name, datatype, d.Value)
			}
		} else if p.asReal(d.Value, &datatype) {
			stmt += fmt.Sprintf("%s %s not null default %f,", d.Name, datatype, d.Value)
		} else if p.asText(d.Value, d.Size, &datatype) {
			stmt += fmt.Sprintf("%s %s not null default '%s',", d.Name, datatype, d.Value)
		} else if p.asBool(d.Value, &datatype) {
			stmt += fmt.Sprintf("%s %s not null default %t,", d.Name, datatype, d.Value)
		} else if p.asTime(d.Value, &datatype) {
			stmt += fmt.Sprintf("%s %s not null default %d,", d.Name, datatype, d.Value.(time.Time).Unix())
		} else {
			return // TODO 2017-02-09: log unrecognized type error
		}
	}
	q := fmt.Sprintf("create table if not exists %s(%s)", table, stmt[:len(stmt)-1])
	log.DB(q, log.LvlDbg)
	if _, err := DB().Exec(q); err != nil {
		log.DB(err, log.LvlErr)
	}
}

func (p *postgres) credentials() string {
	return "postgres://" +
		conf.Instance().DBUser + ":" +
		conf.Instance().DBPass + "@" +
		conf.Instance().DBHost + "/" +
		conf.Instance().DBName + "?" +
		"sslmode=disable"
}

func (p *postgres) asInt(v interface{}, outType *string) bool {
	switch v.(type) {
	case int8, uint8, int16:
		*outType = "smallint"
	case int32, uint16, int:
		*outType = "integer"
	case int64, uint32, uint:
		*outType = "bigint"
	default:
		return false
	}
	return true
}

func (p *postgres) asReal(v interface{}, outType *string) bool {
	switch v.(type) {
	case float32:
		*outType = "real"
	case float64:
		*outType = "double precision"
	default:
		return false
	}
	return true
}

func (p *postgres) asBool(v interface{}, outType *string) bool {
	switch v.(type) {
	case bool:
		*outType = "boolean"
		return true
	}
	return false
}

func (p *postgres) asText(v interface{}, size Datasize, outType *string) bool {
	switch v.(type) {
	case string:
		if size == All {
			*outType = "text"
		} else {
			*outType = "varchar(" + strconv.Itoa(int(size)) + ")"
		}
	case []byte:
		*outType = "bytea"
	default:
		return false
	}
	return true
}

func (p *postgres) asTime(v interface{}, outType *string) bool {
	switch v.(type) {
	case time.Time:
		*outType = "bigint"
		return true
	}
	return false
}
