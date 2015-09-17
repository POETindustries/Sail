package page

import (
	"fmt"
	"html/template"
	"io"
	"sail/conf"
	"sail/dbase"
	"sail/errors"
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

const pagID = "id"
const pagTitle = "title"
const pagContent = "content"
const pagDomain = "domain"
const pagURL = "url"

const pageKeys = pagID + "," + pagTitle + "," + pagContent + "," + pagDomain + "," + pagURL

// Page contains the information needed to generate a web page for display.
// This is the basic struct that contains all information needed to generate
// a correct and complete html page. It is the responsibility of the other
// functions and methods in package page to make sure its fields are
// properly initialized.
type Page struct {
	ID      uint32
	Title   string
	Content template.HTML
	Domain  *Domain
	URL     string

	Meta     *Meta
	template *template.Template

	Conn   *dbase.Conn
	Config *conf.Config
}

// LoadMeta reads metadata from the database and prepares it for display.
// The page's Meta field stores elements like page title, description,
// keywords and other information that is inserted into the html document's
// head area.
func (p *Page) loadMeta() {
	meta := Meta{ID: p.Domain.ID}
	if meta.ScanFromDB(p.Conn) {
		p.Meta = &meta
	} else {
		p.Meta = &Meta{}
	}
}

func (p *Page) loadDomain() {
	if !p.Domain.ScanFromDB(p.Conn) {
		p.Domain = &Domain{ID: 0}
	}
}

func (p *Page) loadTemplate() {
	dir := p.Config.TmplDir + p.Domain.Template

	t, err := template.ParseGlob(dir + "/*.html")
	if err != nil {
		errors.Log(err, p.Config.DevMode)
		t, _ = template.New("frame").Parse(NOTFOUND404)
	}

	p.template = t
}

// ScanFromDB writes data fetched from the database into the
// members of the Page object.
func (p *Page) ScanFromDB(attr string, val interface{}) bool {
	var content string

	p.Domain = &Domain{}
	data := p.Conn.PageData(pageKeys, attr, val)

	if err := data.Scan(&p.ID, &p.Title, &content, &p.Domain.ID, &p.URL); err != nil {
		errors.Log(err, true)
		return false
	}

	p.Content = template.HTML(content)

	return true
}

func (p *Page) Execute(wr io.Writer, data interface{}) error {
	err := p.template.ExecuteTemplate(wr, "frame.html", &data)

	if err != nil {
		errors.Log(err, p.Config.DevMode)
	}

	return err
}

// New creates and returns a Page object. It takes the unique url
// path to the specified page and a database as parameters.
//
// Build always returns a Page object. If there is no page with the given
// name or if there is, but scanning the dataset returns an error, a 404
// page will be returned. Otherwise, the page will be fully constructed
// using its load* methods and a pointer to it is returned.
func New(url string, conn *dbase.Conn, config *conf.Config) *Page {
	p := Page{Conn: conn, Config: config}

	if len(url) <= 1 || !p.ScanFromDB(pagURL, url) {
		if !p.ScanFromDB(pagID, 1) {
			return Load404()
		}
	}

	p.loadDomain()
	fmt.Printf("Domain: %+v\n\n", p.Domain)
	p.loadMeta()
	fmt.Printf("Domain meta data: %+v\n\n", p.Meta)
	p.loadTemplate()
	return &p
}

// Load404 is called whenever generating a page fails somewhere in the process.
// It generates a default error page that informs the user that something
// went wrong when processing their request.
func Load404() *Page {
	tmpl, _ := template.New("frame").Parse(NOTFOUND404)
	p := Page{
		ID:       0,
		Domain:   &Domain{ID: 0},
		Title:    "Sorry about that!",
		Meta:     &Meta{},
		template: tmpl,
		Content:  template.HTML("")}

	// TODO this is the barest minimum of a page. If possible, load another
	// template that fits the corporate design better and generate a 404 out
	// of that one.
	return &p
}
