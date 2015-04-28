// Package conf contains constants and other basic website setup parameters.
// It should serve as a central provider of systemwide used config variables.
// Filepath names, template directories etc.
//
// While some of these variables are
// technically not constants (due to their dependence on the executable's wd),
// they still follow the typical naming convention of writing constants in all
// upper case. This is intentional. It is a strong suggestion to work with them
// as if they were constants, becuase for all intents and purposes, they are.
// Changing their value without exactly knowing what you are doing will have
// severe repercussions.
package conf

import (
	"os"
	"path/filepath"
)

var CWD string

// TMPLDIR is defined relative to the Go executable. This is probably the way
// we'll do it in the future.
var TMPLDIR string

// InitConf basically fills the config variables with values.
func InitConf() {
	// TODO check for errors below
	CWD, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	TMPLDIR = CWD + "/tmpl/"
}
