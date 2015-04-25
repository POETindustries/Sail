// Package page creates html markup from database data and template files.
package page

import (
	"database/sql"
	"html/template"
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
	frame, _ := template.New("frame").Parse(NOTFOUND404)
	p := Page{
		Id:      0,
		Domain:  "error",
		Title:   "Sorry about that!",
		Frame:   frame,
		Meta:    &Meta{},
		Content: template.HTML("")}

	// TODO this is the barest minimum of a page. If possible, load another
	// template that fits the corporate design better and generate a 404 out
	// of that one.

	return &p
}
