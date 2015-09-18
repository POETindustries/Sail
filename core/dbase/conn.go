package dbase

import (
	"database/sql"
	"sail/core/conf"
	"sail/core/errors"

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

// Data should be used for generic queries by objects that use their own
// table information. It is intended for "override" functionality by
// objects that have this database handle as one of their components.
func (c *Conn) Data(table, attrs, attr string, val interface{}) *sql.Row {
	return c.queryRow(table, attrs, attr, val)
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

// New creates a connection object to access the database.
//
// If this fails, it is usually because of wrong credentials or because
// there is no sql service running or the service does not listen on
// the default port.
func New() *Conn {
	conn := &Conn{
		Credentials: conf.Instance().DBCredString(),
		DevMode:     conf.Instance().DevMode}

	if conn.init() == nil && conn.Verify() {
		conn.execCreateInstructs(createInstructs)
		return conn
	}
	return nil
}

func AppendToSchema(schema []string) {
	createInstructs = append(createInstructs, schema...)
}
