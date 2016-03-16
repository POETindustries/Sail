package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sail/errors"
)

const dbUser = "sl_user"
const dbPass = "sl_pass"
const dbName = "sl_main"
const dbHost = "localhost"

const devMode = true
const firstRun = false

type Config struct {
	Cwd       string
	StaticDir string
	TmplDir   string
	ImgDir    string
	JsDir     string
	ThemeDir  string

	DBUser string `json:"db_user"`
	DBPass string `json:"db_password"`
	DBHost string `json:"db_host"`
	DBName string `json:"db_name"`

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
		ImgDir:    cwd + "/static/img/",
		JsDir:     cwd + "/static/js/",
		ThemeDir:  cwd + "/static/theme/"}

	if err := instance.load("config.json"); err != nil {
		errors.Log(err, true)

		instance.DBUser = dbUser
		instance.DBPass = dbPass
		instance.DBHost = dbHost
		instance.DBName = dbName
		instance.DevMode = devMode
		instance.FirstRun = firstRun
	}

	return instance
}

func Instance() *Config {
	if instance == nil {
		new()
	}
	return instance
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
