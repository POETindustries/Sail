package dbase

import "sail/core/plugin"

func Init() {
	plugin.InitDBSchema(schema)
}
