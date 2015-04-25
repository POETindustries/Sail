// Package page creates html markup and templates from database data.
package page

import (
	"database/sql"
	"sail/dbase"
)

func Builder(name string, db *sql.DB) *Page {
	var p Page
	query := "select id,domain,title from sl_page where in_url=?"

	if row := dbase.QueryRow(query, db, name); row == nil {
		return Load404()
	} else if err := row.Scan(&p.Id, &p.Domain, &p.Title); err != nil {
		println(err.Error())
		return Load404()
	} else {
		p.LoadMeta(db)
		p.LoadFrame(db)
		p.LoadContent(db)
	}
	return &p
}

// Load404 is called whenever generating a page fails somewhere in the process.
// It generates a default error page that informs the user that something
// went wrong when processing their request.
func Load404() *Page {
	println("404 loaded")
	return nil
}
