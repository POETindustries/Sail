package errors

import "errors"

// Log allows customized error handling. The error e can be handled
// dependent on the flag devMode, which allows, for example, simple
// output to stdout when in development, and intricate logging and
// webmaster messaging on a per-case basis in production environments.
func Log(e error, devMode bool) {
	if e != nil {
		if devMode {
			println(e.Error())
		} else {
			// TODO more intricate error logging
		}
	}
}

// NoPermission returns a general error that can be raised on
// permission-related issues.
func NoPermission() error {
	return errors.New("ERROR: No permission to access the object")
}

// NoArguments returns a general error that can be raised when
// a function requires values to be set, but finds them empty or not
// present at all.
func NoArguments() error {
	return errors.New("ERROR: No arguments passed")
}

// NilPointer returns a general error that signifies an attempt to
// dereference a nil pointer.
func NilPointer() error {
	return errors.New("ERROR: Nil pointer dereferenced")
}
