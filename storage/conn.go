package storage

import (
	"database/sql"
	"sail/conf"
	"sail/errors"

	_ "github.com/lib/pq" // driver for postgres
)

type Conn struct {
	DB          *sql.DB
	Credentials string
	DevMode     bool
}

var instance *Conn

// New creates a connection object to access the database.
//
// If this fails, it is usually because of wrong credentials or because
// there is no sql service running or the service does not listen on
// the default port.
func new() *Conn {
	instance = &Conn{
		Credentials: conf.Instance().DBCredString(),
		DevMode:     conf.Instance().DevMode}

	if instance.init() == nil && instance.Verify() {
		instance.execCreateInstructs(createInstructs)
		return instance
	}
	return nil
}

func Instance() *Conn {
	if instance == nil {
		new()
	}
	return instance
}

func (c *Conn) Verify() bool {
	if err := c.DB.Ping(); err != nil {
		errors.Log(err, c.DevMode)
		return false
	}
	return true
}

func (c *Conn) init() (err error) {
	c.DB, err = sql.Open("postgres", c.Credentials)

	if err != nil {
		errors.Log(err, c.DevMode)
	}
	return
}

func (c *Conn) PageData(attr string, val interface{}) *sql.Row {
	return c.queryRow("sl_page", PageAttrs, attr, val)
}

func (c *Conn) DomainData(id uint32) *sql.Row {
	return c.queryRow("sl_domain", DomainAttrs, domainID, id)
}

func (c *Conn) queryRow(table, proj, attr string, val interface{}) *sql.Row {
	sel := "where " + attr + " =$1"
	query := "select " + proj + " from " + table + " " + sel

	return c.DB.QueryRow(query, val)
}

func (c *Conn) execCreateInstructs(instructs []string) (err error) {
	for _, instruct := range instructs {
		if _, err = c.DB.Exec(instruct); err != nil {
			errors.Log(err, conf.Instance().DevMode)
		}
	}
	return
}
