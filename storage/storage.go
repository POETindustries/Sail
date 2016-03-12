package storage

import (
	"database/sql"
	"os"
	"sail/conf"
	"sail/errors"
	pageschema "sail/page/schema"
	userschema "sail/user/schema"

	// sqlite database driver
	_ "github.com/mattn/go-sqlite3"
)

const version = 1

var db *sql.DB
var createInstructs = []string{
	userschema.CreateUser,
	userschema.InitUser,
	pageschema.CreateWidget,
	pageschema.InitWidget,
	pageschema.CreateWidgetMenu,
	pageschema.InitWidgetMenu,
	pageschema.CreateWidgetText,
	pageschema.CreateTemplate,
	pageschema.InitTemplate,
	pageschema.CreateTemplateWidgets,
	pageschema.InitTemplateWidgets,
	pageschema.CreateMeta,
	pageschema.InitMeta,
	pageschema.CreatePage,
	pageschema.InitPage}

// DB returns a pointer to the database handle singleton.
func DB() *sql.DB {
	if db == nil && !dbInit() {
		panic("storage: Database init failed")
	}
	return db
}

func dbInit() bool {
	loc := "db/"
	if _, err := os.Stat(loc); err != nil {
		os.MkdirAll(loc, 0700)
	}
	db, _ = sql.Open("sqlite3", loc+"panoptiq.db")
	return db.Ping() == nil
}

// ExecCreateInstructs takes care of first-time setup of the datastore.
func ExecCreateInstructs() (err error) {
	for _, instruct := range createInstructs {
		if _, err = DB().Exec(instruct); err != nil {
			errors.Log(err, conf.Instance().DevMode)
		}
	}
	return
}
