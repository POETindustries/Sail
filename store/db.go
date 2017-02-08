package store

import (
	"database/sql"
	"sail/conf"
)

// Driver provides unified access to the actual sql drivers.
type Driver interface {
	Copy() Driver
	Init() (*sql.DB, error)
	Data(query *Query) []interface{}
	Setup(data *SetupData)
	credentials() string
}

// SetupData provides a way to send all necessary data to
// the Database's Setup function.
type SetupData struct {
	Data     map[string]interface{}
	Primary  string
	Relation string
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
		if !instance.init() {
			panic("Database initialization failed!")
		}
	}
	return instance
}

// Exec wraps the function of the same name from sql.DB.
func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}

// Ping wraps the function of the same name from sql.DB.
func (d *Database) Ping() error {
	return d.db.Ping()
}

// Query wraps the function of the same name from sql.DB.
func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.Query(query, args...)
}

// QueryRow wraps the function of the same name from sql.DB.
func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.db.QueryRow(query, args...)
}

// Setup uses the data passed to build creation queries that
// fit the underlying SQL dialect. Packages using the store
// package can provide data to be stored without having to
// know about storage internals.
func (d *Database) Setup(data *SetupData) {
	d.driver.Setup(data)
}

// init sets up proper initialization of the database backend.
// Failure indicates a serious problem and should prevent the
// app from continuing.
func (d *Database) init() bool {
	switch conf.Instance().DBDriver {
	case "sqlite3":
		d.driver = &sqlite3{}
	case "mysql":
		d.driver = &mysql{}
	case "postgres":
		d.driver = &postgres{}
	default:
		return false
	}
	if db, err := d.driver.Init(); err == nil {
		d.db = db
		return (d.db.Ping() == nil)
	}
	return false
}
