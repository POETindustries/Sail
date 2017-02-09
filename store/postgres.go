package store

import (
	"database/sql"
	"sail/conf"
	"strconv"
	"strings"

	// postgres database driver
	_ "github.com/lib/pq"
)

type postgres struct {
	count int
}

func (p *postgres) Copy() Driver {
	return &postgres{}
}

func (p *postgres) Data(query *Query) []interface{} {
	for i, a := range query.attrs {
		if strings.HasSuffix(a, "?") {
			p.count++
			query.attrs[i] = strings.Replace(a, "?", "$"+strconv.Itoa(p.count), 1)
		}
	}
	for i, s := range query.selection {
		if strings.HasSuffix(s, "?") {
			p.count++
			query.selection[i] = strings.Replace(s, "?", "$"+strconv.Itoa(p.count), 1)
		}
	}
	return append(query.attrVals, query.selectionVals...)
}

func (p *postgres) Init() (*sql.DB, error) {
	return sql.Open("postgres", p.credentials())
}

func (p *postgres) Setup(table string, data []*SetupData) {

}

func (p *postgres) credentials() string {
	return "postgres://" +
		conf.Instance().DBUser + ":" +
		conf.Instance().DBPass + "@" +
		conf.Instance().DBHost + "/" +
		conf.Instance().DBName + "?" +
		"sslmode=disable"
}
