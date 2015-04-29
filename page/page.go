package page

import (
	"database/sql"
	"html/template"
	"sail/conf"
	"sail/dbase"
	"sail/tmpl"
	"strings"
)

// NOTFOUND404 is a very basic web page signaling a 404 error.
// It contails the bare minimum necessary for a syntactically correct html web
// page and is used in those cases when not even basic database connections
// and templates work. The cms cannot be considered functional should that
// happen, and this markup at least tells the user as much. The markup is as
// generic as possible while still being somewhat good looking.
const NOTFOUND404 = `<!doctype html>
		<html style="background:black;text-align:center;color:white;">
		<head><title>Sorry About That</title><meta charset="utf-8"></head>
		<body style="padding:72px;font-family:sans-serif;font-size:1.5em;">
		<p style="font-size:2em;">Sorry About That!</p>
		<p>PAGE NOT FOUND</p></body></html>`

// Page contains the information needed to generate a web page for display.
// This is the basic struct that contains all information needed to generate
// a correct and complete html page. It is the responsibility of the other
// functions and methods in package page to make sure its fields are
// properly initialized.
type Page struct {
	ID      int32
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
	var csv string
	templates := []string{conf.TMPLDIR + "404.html"}
	query := "select frame_tmpl from sl_page_meta where domain=?"

	if row := dbase.QueryRow(query, db, p.Domain); row != nil {
		if row.Scan(&csv) == nil {
			templates = tmpl.PrepareFiles(strings.Split(csv, ","))
		}
	} //else: templates still contains only the 404 page

	if p.Frame, err = template.ParseFiles(templates...); err != nil {
		println(err.Error())
		p.Frame, _ = template.New("frame").Parse(NOTFOUND404)
	}
}

// LoadContent fetches the page's content from the database. Content is that
// piece of a web page that is usually generated in the backend by someone
// working with the cms. It should allow some subsets of html tags, but it
// is usually a good idea to sanitize JavaScript.
func (p *Page) LoadContent(db *sql.DB) {
	var content string
	query := "select content from sl_page where id=?"

	if row := dbase.QueryRow(query, db, p.ID); row != nil {
		if err := row.Scan(&content); err != nil {
			println(err.Error())
		}
	} // else: content still has zero value and is just an empty string
	p.Content = template.HTML(content)
}
