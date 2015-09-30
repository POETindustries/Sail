package page

import (
	"fmt"
	"html/template"
	"io"
	"sail/conf"
	"sail/domain"
	"sail/errors"
	"sail/storage"
)

// Page contains the information needed to generate a web page for display.
// This is the basic struct that contains all information needed to generate
// a correct and complete html page. It is the responsibility of the other
// functions and methods in package page to make sure its fields are
// properly initialized.
type Page struct {
	ID      uint32
	Title   string
	URL     string
	Content template.HTML
	Domain  *domain.Domain

	Status int8
	Owner  string
	CDate  string
	EDate  string

	conn   *storage.Conn
	config *conf.Config
}

func (p *Page) loadDomain() {
	if !p.Domain.ScanFromDB() {
		p.Domain = &domain.Domain{ID: 0}
	}
}

// ScanFromDB writes data fetched from the database into the
// members of the Page object.
func (p *Page) ScanFromDB(attr string, val interface{}) bool {
	content := ""
	dom := domain.Domain{}
	data := p.conn.PageData(attr, val)

	if err := data.Scan(&p.ID,
		&p.Title,
		&content,
		&dom.ID,
		&p.URL,
		&p.Status,
		&p.Owner,
		&p.CDate,
		&p.EDate); err != nil {
		errors.Log(err, true)
		return false
	}

	p.Content = template.HTML(content)
	p.Domain = &dom

	return true
}

func (p *Page) Execute(wr io.Writer) error {
	return p.Domain.Template.Execute(wr, p)
}

// New creates and returns a Page object. It takes the unique url
// path to the specified page and a database as parameters.
//
// Build always returns a Page object. If there is no page with the given
// name or if there is, but scanning the dataset returns an error, a 404
// page will be returned. Otherwise, the page will be fully constructed
// using its load* methods and a pointer to it is returned.
func New(url string) *Page {
	p := Page{conn: storage.Instance(), config: conf.Instance()}

	if len(url) <= 1 || !p.ScanFromDB(storage.PageURL, url) {
		if !p.ScanFromDB(storage.PageID, 1) {
			return Load404()
		}
	}

	p.loadDomain()
	fmt.Printf("Domain: %+v\n\n", p.Domain)

	return &p
}

// Load404 is called whenever generating a page fails somewhere in the process.
// It generates a default error page that informs the user that something
// went wrong when processing their request.
func Load404() *Page {
	p := Page{
		ID:      0,
		Domain:  &domain.Domain{ID: 0, Meta: &domain.Meta{}},
		Title:   "Sorry about that!",
		Content: template.HTML("")}

	// TODO this is the barest minimum of a page. If possible, load another
	// template that fits the corporate design better and generate a 404 out
	// of that one.
	return &p
}
