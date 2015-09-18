package plugin

import (
	"sail/core/conf"
	"sail/core/dbase"
)

func ConfigData() *conf.Config {
	return conf.Instance()
}

func DBConnection() *dbase.Conn {
	return dbase.New()
}

func InitDBSchema(schema []string) {
	dbase.AppendToSchema(schema)
}
