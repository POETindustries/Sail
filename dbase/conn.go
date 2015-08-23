package dbase

import (
	"database/sql"
	"sail/errors"
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
