package storage

import (
	"sail/conf"
	"sail/errors"
	"sail/storage/psqldb"
	"sail/storage/schema"
)

// ExecCreateInstructs takes care of first-time setup of the datastore.
func ExecCreateInstructs() (err error) {
	for _, instruct := range schema.CreateInstructs {
		if _, err = psqldb.Instance().DB.Exec(instruct); err != nil {
			errors.Log(err, conf.Instance().DevMode)
		}
	}
	return
}
