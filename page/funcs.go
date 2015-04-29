// Package page creates html markup from database data and template files.
//
// Trustworthiness of Data
//
// The contract for handling database content is that methods and functions
// that read from the database assume that the database content is trustworthy
// while functions that write to the database assume that user input is not
// to be trusted.
//
// This is only safe as long as an intruder does not manage to write harmful
// code directly into the database by bypassing the filters used by write
// functions. This needs to be addressed.
//
// On Separation Logic From Views
//
// It is considered good practice to separate data, logic and outside views
// as much as possible (hence, MVC patterns and such). This package is
// something of an exception to this otherwise reasonable rule. There are
// some string constant and functions that contain html markup directly
// embedded into the code.
//
// The reason is simple: If all else fails, we still want to be able to let
// the user know that there is a problem with the website and that they should
// consider coming back later. We cannot load templates if something with
// loading templates is wrong and we cannot load data from databases if
// loading from databases is broken, so we have to assume that in the worst
// case scenario nothing else works other than simplest code.
//
// This is why there is some hardcoded html in this package, acting as some
// kind of failsafe.
package page

import (
	"database/sql"
	"html/template"
	"sail/dbase"
)

// Builder creates and returns a Page object. It takes the unique url
// path to the specified page and a database as parameters.
//
// Builder always returns a Page object. If there is no page with the given
// name or if there is, but scanning the dataset returns an error, a 404
// page will be returned. Otherwise, the page will be fully constructed
// using its load* methods and a pointer to it is returned.
func Builder(inUrl string, db *sql.DB) *Page {
	var p Page
	query := "select id,domain,title from sl_page where in_url=?"

	if row := dbase.QueryRow(query, db, inUrl); row == nil {
		return Load404()
	} else if err := row.Scan(&p.ID, &p.Domain, &p.Title); err != nil {
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
		ID:      0,
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
