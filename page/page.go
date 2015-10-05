package page

import (
	"html/template"
	"sail/domain"
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
	Content string
	Domain  *domain.Domain

	Status int8
	Owner  string
	CDate  string
	EDate  string
}

// Markup wraps p.Content into an html-friendly way, ready for
// usage in a template.
func (p *Page) Markup() template.HTML {
	return template.HTML(p.Content)
}

// New creates a new Page object with usable defaults.
func New() *Page {
	return &Page{Domain: domain.New()}
}
