// Trustworthiness of data
//
// The contract for handling database content is that methods and functions
// that read from the database assume that the database content is trustworthy
// while functions that write to the database assume that user input is not
// to be trusted.
//
// This is only safe as long as an intruder does not manage to write harmful
// code directly into the database by bypassing the filters used by write
// functions. This needs to be addressed.
package page

import (
	"database/sql"
	"html/template"
	"sail/conf"
	"sail/dbase"
)

// Page contains the information needed to generate a web page for display.
type Page struct {
	Id      int32
	Domain  string
	Title   string
	Meta    *Meta
	Frame   *template.Template
	Content template.HTML
}

// LoadMeta reads metadata from the database and prepares it for display.
// The page's Meta field stores elements like page title, description,
// keywords and other information that is inserted into the html document's
// head area.
func (p *Page) LoadMeta(db *sql.DB) {
	var meta Meta
	query := "select " + DBMETAKEYS + " from sl_page_meta where domain=?"
	if row := dbase.QueryRow(query, db, p.Domain); row != nil {
		row.Scan(&meta.Title,
			&meta.Keywords,
			&meta.Description,
			&meta.Language,
			&meta.PageTopic,
			&meta.RevisitAfter,
			&meta.Robots)
	} //else: meta has zero values, will be full of empty strings
	p.Meta = &meta
}

// LoadFrame fetches the html template file that belongs to the page's domain.
// It generates a template object to be passed to Page.Frame where it can be
// fetched later to generate the whole html page.
func (p *Page) LoadFrame(db *sql.DB) {
	var err error
	templateFile := "404.html"
	query := "select frame_tmpl from sl_page_meta where domain=?"

	if row := dbase.QueryRow(query, db, p.Domain); row != nil {
		if err = row.Scan(&templateFile); err == nil {
			templateFile += ".html"
		}
	} //else: templateFile points to a default 404 page
	if p.Frame, err = template.ParseFiles(conf.DOCROOT + templateFile); err != nil {
		println(err.Error())
	}
}

// LoadContent fetches the page's content from the database. Content is that
// piece of a web page that is usually generated in the backend by someone
// working with the cms.
func (p *Page) LoadContent(db *sql.DB) {
	var content string
	query := "select content from sl_page where id=?"

	if row := dbase.QueryRow(query, db, p.Id); row != nil {
		if err := row.Scan(&content); err != nil {
			println(err.Error())
		}
	} // else: content still has zero value and is just an empty string
	p.Content = template.HTML(content)
}

// Load404 is called whenever generating a page fails somewhere in the process.
// It generates a default error page that informs the user that something
// went wrong when processing their request.
func (p *Page) load404() {
	println("404 loaded")
}
