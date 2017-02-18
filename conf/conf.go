package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sail/errors"
)

const skeleton = `{
"db_driver"     : "%s",
"db_user"       : "%s",
"db_password"   : "%s",
"db_name"       : "%s",
"db_host"       : "%s",

"dev_mode"      : %t,
"first_run"     : %t
}`

const (
	dbDriver = "sqlite3"
	dbUser   = "sl_user"
	dbPass   = "sl_pass"
	dbName   = "sl_main"
	dbHost   = "localhost"

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

	DevMode  bool `json:"dev_mode"`
	FirstRun bool `json:"first_run"`
}

var instance *Config

func new() *Config {
	cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	instance = &Config{
		Cwd:       cwd + "/",
		StaticDir: cwd + "/static/",
		TmplDir:   cwd + "/static/tmpl/",
		FileDir:   cwd + "/static/files/",
		JsDir:     cwd + "/static/js/",
		ThemeDir:  cwd + "/static/theme/"}

	if err := instance.load("config.json"); err != nil {
		errors.Log(err, true)
		instance.DBDriver = dbDriver
		instance.DBUser = dbUser
		instance.DBPass = dbPass
		instance.DBHost = dbHost
		instance.DBName = dbName
		instance.DevMode = devMode
		instance.FirstRun = firstRun
	}
	return instance
}

// Instance provides access to the application-wide config
// singleton.
func Instance() *Config {
	if instance == nil {
		new()
	}
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
		c.DBHost, c.DevMode, c.FirstRun)
	err := ioutil.WriteFile(c.Cwd+"config.json", []byte(f), 0640) //maybe do 0600 for hardened security?
	if err != nil {
		errors.Log(err, c.DevMode)
	}
}
