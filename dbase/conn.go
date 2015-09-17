package dbase

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

func (c *Conn) PageData(attrs, attr string, val interface{}) *sql.Row {
	return c.queryRow("sl_page", attrs, attr, val)
}

func (c *Conn) MetaData(attrs string, id uint32) *sql.Row {
	return c.queryRow("sl_page_meta", attrs, "id", id)
}

func (c *Conn) DomainData(attrs string, id uint32) *sql.Row {
	return c.queryRow("sl_domain", attrs, "id", id)
}

func (c *Conn) queryRow(table, proj, attr string, val interface{}) *sql.Row {
	sel := "where " + attr + " =$1"
	query := "select " + proj + " from " + table + " " + sel

	return c.DB.QueryRow(query, val)
}

// New creates a connection object to access the database.
//
// If this fails, it is usually because of wrong credentials or because
// there is no sql service running or the service does not listen on
// the default port.
func New(config *conf.Config) *Conn {
	conn := &Conn{
		Credentials: config.DBCredString(),
		DevMode:     config.DevMode}

	if conn.init() == nil && conn.Verify() {
		if config.FirstRun {
			for _, instruct := range createInstructs {
				if _, err := conn.DB.Exec(instruct); err != nil {
					errors.Log(err, config.DevMode)
				}
			}
		}
		return conn
	}
	return nil
}
