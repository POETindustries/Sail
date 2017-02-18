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
func NoArguments(what string) error {
	return errors.New("ERROR: No arguments passed - " + what)
}

// NilPointer returns a general error that signifies an attempt to
// dereference a nil pointer.
func NilPointer() error {
	return errors.New("ERROR: Nil pointer dereferenced")
}
