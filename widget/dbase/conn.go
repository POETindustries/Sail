package dbase

import (
	"database/sql"
	"sail/core/plugin"
)

type Conn struct {
}

func Init() {
	plugin.InitDBSchema(schema)
}

func (c *Conn) WidgetData(attrs, attr string, val interface{}) *sql.Row {
	return plugin.DBConnection().Data("sl_widget", attrs, attr, val)
}

func (c *Conn) MenuData(attrs string, id uint32) *sql.Row {
	return plugin.DBConnection().Data("sl_widget_menu", attrs, "id", id)
}

func (c *Conn) StoreMenu() {
	// TODO needs parameters
}

func (c *Conn) DeleteMenu(id uint32) {

}

func (c *Conn) deleteWidgetEntry(id uint32) {

}
