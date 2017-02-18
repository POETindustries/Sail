package store

import (
	"database/sql"
	"sail/conf"
	"sail/errors"
)

type Datasize int16

const (
	All   Datasize = 0
	Small Datasize = 32
	Mid   Datasize = 8192
	Large Datasize = 16384
)

// Driver provides unified access to the actual sql drivers.
type Driver interface {
	Copy() Driver
	Init() (*sql.DB, error)
	Data(query *Query) []interface{}
	Prepare(query string) string
	Setup(table string, data []*SetupData)
	credentials() string
}

// SetupData provides a way to send all necessary data to
// the Database's Setup function.
type SetupData struct {
	Name      string
	Value     interface{}
	IsPrimary bool
	Size      Datasize
}

// Database gives access to the database implementation
// responsible for actually doing all the database work.
type Database struct {
	db     *sql.DB
	driver Driver
}

var instance *Database

// DB provides access to the Database singleton.
func DB() *Database {
	if instance == nil {
		instance = &Database{}
		if err := instance.init(); err != nil {
			errors.Log(err, conf.Instance().DevMode)
			panic("Database initialization failed!")
		}
	}
	return instance
}

// Exec wraps the function of the same name from sql.DB.
func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(d.driver.Prepare(query), args...)
}

// Ping wraps the function of the same name from sql.DB.
func (d *Database) Ping() error {
	return d.db.Ping()
}

// Query wraps the function of the same name from sql.DB.
func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.Query(d.driver.Prepare(query), args...)
}

// QueryRow wraps the function of the same name from sql.DB.
func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.db.QueryRow(d.driver.Prepare(query), args...)
}

// Setup uses the data passed to build creation queries that
// fit the underlying SQL dialect. Packages using the store
// package can provide data to be stored without having to
// know about storage internals.
func (d *Database) Setup(table string, data []*SetupData) {
	d.driver.Setup(table, data)
}

// init sets up proper initialization of the database backend.
// Failure indicates a serious problem and should prevent the
// app from continuing.
func (d *Database) init() error {
	switch conf.Instance().DBDriver {
	case "sqlite3":
		d.driver = &sqlite3{}
	case "mysql":
		d.driver = &mysql{}
	case "postgres":
		d.driver = &postgres{}
	default:
		return ErrNoDriver{conf.Instance().DBDriver}
	}
	db, err := d.driver.Init()
	if err != nil {
		return err
	}
	d.db = db
	return d.db.Ping()
}
