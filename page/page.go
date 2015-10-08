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
	Content template.HTML
	Domain  *domain.Domain

	Status int8
	Owner  string
	CDate  string
	EDate  string
}

// New creates a new Page object with usable defaults.
func New() *Page {
	return &Page{Domain: domain.New()}
}
