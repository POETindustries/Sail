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

package log

import (
	"fmt"
	"os"
	"time"
)

// Loggable can be implemented by objects that should have finer
// control over the output format when logging.
type Loggable interface {
	String() string
}

// Log levels range from no logging at all to generous spilling around.
const (
	// Don't log anything. The app is completely silent.
	lvlNone = 0x00

	// LvlErr should be passed to a log message if it was triggered by an error.
	LvlErr = 0x01

	// LvlWarn should be passed along a log message if it was triggered by behavior
	// that could be harmless, but might cause problems under certain conditions.
	LvlWarn = 0x02

	// LvlDbg can be passed to a log message if it just logs useful information
	// during development and such.
	LvlDbg = 0x04
)

// File descriptors for the different module-specific log files
const (
	loc = "/tmp/sail/"
	db  = loc + "db.log"
	acc = loc + "acc.log"
	srv = loc + "srv.log"
)

// level holds the current log level
var level = lvlNone

// Init prepares the filesystem for logging. It creates the log
// directory if it doesn't exist.
func Init() {
	if _, err := os.Stat(loc); err != nil {
		os.MkdirAll(loc, 0755)
	}
}

// SetLevel can be used to set the current log level. It is designed to
// be set via command line arguments at the start of the program, but
// can be used from anywhere within the code.
func SetLevel(lvl string) bool {
	switch lvl {
	case "error":
		level = LvlErr
	case "warning":
		level = LvlWarn
	case "debug":
		level = LvlDbg
	case "none":
		level = lvlNone
	default:
		return false
	}
	return true
}

// DB logs to the log file reserved for database events.
func DB(obj interface{}, lvl int) {
	log(obj, lvl, db)
}

// Acc logs to the file reserved for events relating to user accounts.
func Acc(obj interface{}, lvl int) {
	log(obj, lvl, acc)
}

// Srv logs all server related events.
func Srv(obj interface{}, lvl int) {
	log(obj, lvl, srv)
}

func log(obj interface{}, lvl int, file string) {
	if level != lvlNone && lvl <= level {
		t := time.Now().String()
		flags := os.O_APPEND | os.O_WRONLY | os.O_CREATE
		f, _ := os.OpenFile(file, flags, 0600)
		switch lvl {
		case LvlErr:
			f.WriteString(t + ":   Error: " + msg(obj) + "\n")
		case LvlWarn:
			f.WriteString(t + ": Warning: " + msg(obj) + "\n")
		case LvlDbg:
			f.WriteString(t + ":   Debug: " + msg(obj) + "\n")
		}
		f.Close()
	}
}

func msg(obj interface{}) string {
	switch obj := obj.(type) {
	case string:
		return obj
	case Loggable:
		return fmt.Sprintf("%s", obj)
	case error:
		return obj.Error()
	default:
		return fmt.Sprintf("%+v", obj)
	}
}
