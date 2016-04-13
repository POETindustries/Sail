package storage

import (
	"database/sql"
	"os"
	"sail/conf"
	"sail/errors"
	fileschema "sail/file/schema"
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
	userschema.CreateGroup,
	userschema.InitGroup,
	userschema.CreateGroupMembers,
	userschema.InitGroupMembers,
	fileschema.CreateFile,
	pageschema.CreateWidget,
	pageschema.InitWidget,
	pageschema.CreateWidgetNav,
	pageschema.InitWidgetNav,
	pageschema.CreateWidgetText,
	pageschema.CreateTemplate,
	pageschema.InitTemplate,
	pageschema.CreateTemplateWidgets,
	pageschema.InitTemplateWidgets,
	pageschema.CreateMeta,
	pageschema.InitMeta,
	pageschema.CreateContent,
	pageschema.InitContent}

// DB returns a pointer to the database handle singleton.
func DB() *sql.DB {
	if db == nil && !dbInit() {
		panic("storage: Database init failed")
	}
	return db
}

// ExecCreateInstructs takes care of first-time setup of the datastore.
func ExecCreateInstructs() (err error) {
	if conf.Instance().FirstRun {
		for _, instruct := range createInstructs {
			if _, err = DB().Exec(instruct); err != nil {
				errors.Log(err, conf.Instance().DevMode)
			}
		}
	}
	return
}

func dbInit() bool {
	loc := conf.Instance().Cwd + "db/"
	if _, err := os.Stat(loc); err != nil {
		os.MkdirAll(loc, 0700)
	}
	db, _ = sql.Open("sqlite3", loc+"sail.db")
	return db.Ping() == nil
}
