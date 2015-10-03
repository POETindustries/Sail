package conn

import (
	"database/sql"
	"sail/conf"
	"sail/errors"

	_ "github.com/lib/pq" // driver for postgres
)

// Conn is the primary data store connection agent.
//
// It handles user credentials as well as actually connecting to and
// pinning the database.
type Conn struct {
	DB          *sql.DB
	credentials string
}

var instance *Conn

// Instance returna a reference to the singleton connection object
// responsible for access to the database.
//
// It returns nil if the first-time creation of the instance object
// failed for any reason, usually because of wrong credentials or
// because there is no sql service running or the service does not
// listen on the default port.
func Instance() *Conn {
	if instance == nil {
		instance = &Conn{credentials: conf.Instance().DBCredString()}
		if instance.init() != nil || !instance.Verify() {
			return nil
		}
	}
	return instance
}

// Verify returns TRUE if the Conn object is able to actually connect
// to the database with the currently stored credentials.
//
// It is automatically called during Conn instance creation through
// the Instance() function, so it does not need to be called separately
// when creating the object.
//
// Verify is a maintenance method and thus can become useful after long
// pauses in web traffic and specifically database access, but it should
// not be necessary to call this method during normal operations,
// especially if traffic is high anyway.
func (c *Conn) Verify() bool {
	if err := c.DB.Ping(); err != nil {
		errors.Log(err, conf.Instance().DevMode)
		return false
	}
	return true
}

func (c *Conn) init() (err error) {
	c.DB, err = sql.Open("postgres", c.credentials)
	if err != nil {
		errors.Log(err, conf.Instance().DevMode)
	}
	return
}
