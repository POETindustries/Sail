package store

import (
	"database/sql"
	"sail/conf"
	"strconv"

	// postgres database driver
	_ "github.com/lib/pq"
)

type postgres struct {
	count int
}

func (p *postgres) Copy() Driver {
	return &postgres{}
}

func (p *postgres) Init() (*sql.DB, error) {
	return sql.Open("postgres", p.credentials())
}

func (p *postgres) Param() string {
	p.count++
	return "$" + strconv.Itoa(p.count)
}

func (p *postgres) credentials() string {
	return "postgres://" +
		conf.Instance().DBUser + ":" +
		conf.Instance().DBPass + "@" +
		conf.Instance().DBHost + "/" +
		conf.Instance().DBName + "?" +
		"sslmode=disable"
}
