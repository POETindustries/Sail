package errors

import "errors"

func Log(e error, devMode bool) {
	if e != nil {
		if devMode {
			println(e.Error())
		} else {
			// TODO more intricate error logging
		}
	}
}

func NewNoPermission() error {
	return errors.New("No permission to access the object")
}
