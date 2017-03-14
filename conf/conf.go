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

package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sail/log"
	"sync"
)

const skeleton = `{
"db_driver"     : "%s",
"db_user"       : "%s",
"db_password"   : "%s",
"db_name"       : "%s",
"db_host"       : "%s",

"mail_user"      : "%s",
"mail_pass"      : "%s",
"mail_address"   : "%s",
"mail_host_smtp" : "%s",
"mail_port_smtp" : %d,

"dev_mode"      : %t,
"first_run"     : %t
}
`

const (
	dbDriver = "sqlite3"
	dbUser   = "sl_user"
	dbPass   = "sl_pass"
	dbName   = "sl_main"
	dbHost   = "localhost"

	mailUser     = "Sail"
	mailPortSMTP = 587

	devMode  = true
	firstRun = false
)

// Config holds the application's basic configuration. It
// is read from a config file that server administrators have
// access to.
type Config struct {
	Cwd       string `json:"-"`
	StaticDir string `json:"-"`
	TmplDir   string `json:"-"`
	FileDir   string `json:"-"`
	JsDir     string `json:"-"`
	ThemeDir  string `json:"-"`

	DBDriver string `json:"db_driver"`
	DBUser   string `json:"db_user"`
	DBPass   string `json:"db_password"`
	DBHost   string `json:"db_host"`
	DBName   string `json:"db_name"`

	MailUser     string `json:"mail_user"`
	MailPass     string `json:"mail_pass"`
	MailAddress  string `json:"mail_address"`
	MailHostSMTP string `json:"mail_host_smtp"`
	MailPortSMTP uint16 `json:"mail_port_smtp"`

	DevMode  bool `json:"dev_mode"`
	FirstRun bool `json:"first_run"`
}

var instance *Config
var initializer sync.Once

func new() *Config {
	cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	c := &Config{
		Cwd:       cwd + "/",
		StaticDir: cwd + "/static/",
		TmplDir:   cwd + "/static/tmpl/",
		FileDir:   cwd + "/static/files/",
		JsDir:     cwd + "/static/js/",
		ThemeDir:  cwd + "/static/theme/"}

	if err := c.load("config.json"); err != nil {
		log.Srv(err, log.LvlWarn)
		c.DBDriver = dbDriver
		c.DBUser = dbUser
		c.DBPass = dbPass
		c.DBHost = dbHost
		c.DBName = dbName
		c.MailUser = mailUser
		c.MailPortSMTP = mailPortSMTP
		c.DevMode = devMode
		c.FirstRun = firstRun
	}
	return c
}

// Instance provides access to the application-wide config
// singleton.
func Instance() *Config {
	initializer.Do(func() {
		instance = new()
	})
	return instance
}

func (c *Config) load(file string) error {
	in, err := ioutil.ReadFile(c.Cwd + file)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(in, c); err != nil {
		return err
	}
	return nil
}

// Save writes the current state of the Config object to the
// config file.
func (c *Config) Save() {
	f := fmt.Sprintf(skeleton, c.DBDriver, c.DBUser, c.DBPass, c.DBName,
		c.DBHost, c.MailUser, c.MailPass, c.MailAddress, c.MailHostSMTP,
		c.MailPortSMTP, c.DevMode, c.FirstRun)
	err := ioutil.WriteFile(c.Cwd+"config.json", []byte(f), 0640) //maybe do 0600 for hardened security?
	if err != nil {
		log.Srv(err, log.LvlWarn)
	}
}
