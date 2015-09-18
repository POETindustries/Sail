package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sail/core/errors"
)

const dbUser = "sl_user"
const dbPass = "sl_pass"
const dbName = "sl_main"
const dbHost = "localhost"

const devMode = true
const firstRun = false

type Config struct {
	Cwd      string
	TmplDir  string
	ImgDir   string
	JsDir    string
	ThemeDir string

	DBUser string `json:"db_user"`
	DBPass string `json:"db_password"`
	DBHost string `json:"db_host"`
	DBName string `json:"db_name"`

	DevMode  bool `json:"dev_mode"`
	FirstRun bool `json:"first_run"`
}

func New() *Config {
	cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	config := Config{
		Cwd:      cwd + "/",
		TmplDir:  cwd + "/tmpl/",
		ImgDir:   cwd + "/img/",
		JsDir:    cwd + "/js/",
		ThemeDir: cwd + "/theme/"}

	if err := config.load("development.conf"); err != nil {
		errors.Log(err, true)

		config.DBUser = dbUser
		config.DBPass = dbPass
		config.DBHost = dbHost
		config.DBName = dbName
		config.DevMode = devMode
		config.FirstRun = firstRun
	}

	return &config
}

func (c *Config) DBCredString() string {
	return "postgres://" +
		c.DBUser + ":" +
		c.DBPass + "@" +
		c.DBHost + "/" +
		c.DBName + "?" +
		"sslmode=disable"
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
