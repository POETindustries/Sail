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
	"sail/conf"
	"sail/log"
	"sync"
)

type Datasize int16

const (
	All   Datasize = 0
	Small Datasize = 64
	Mid   Datasize = 256
	Large Datasize = 4096
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
var initializer sync.Once

// DB provides access to the Database singleton.
func DB() *Database {
	initializer.Do(func() {
		instance = &Database{}
		if err := instance.init(); err != nil {
			log.DB(err, log.LvlErr)
			panic("Database initialization failed!")
		}
	})
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
