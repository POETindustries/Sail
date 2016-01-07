package storage

import (
	"sail/conf"
	"sail/errors"
	"sail/storage/domainstore"
	"sail/storage/pagestore"
	"sail/storage/psqldb"
	"sail/storage/templatestore"
	"sail/storage/userstore"
	"sail/storage/widgetstore"
)

var createInstructs = [][]string{
	widgetstore.CreateInstructs,
	templatestore.CreateInstructs,
	domainstore.CreateInstructs,
	pagestore.CreateInstructs,
	userstore.CreateInstructs}

// ExecCreateInstructs takes care of first-time setup of the datastore.
func ExecCreateInstructs() (err error) {
	for _, instructs := range createInstructs {
		for _, instruct := range instructs {
			if _, err = psqldb.Instance().DB.Exec(instruct); err != nil {
				errors.Log(err, conf.Instance().DevMode)
			}
		}
	}
	return
}
