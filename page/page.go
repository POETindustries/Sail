package page

import "sail/tmpl"

// Page contains the information needed to generate a web page for display.
// This is the basic struct that contains all information needed to generate
// a correct and complete html page. It is the responsibility of the other
// functions and methods in package page to make sure its fields are
// properly initialized.
type Page struct {
	ID       uint32
	Title    string
	URL      string
	Content  string
	Meta     *Meta
	Template *tmpl.Template

	Status int8
	Owner  string
	CDate  string
	EDate  string
}

// Meta holds the meta information of a web page. It is used to store values
// for display in an html page's <head> block. This struct holds values that
// are used foremost for SEO purposes. Some meta information that is not page
// specific and doesn't really change across websites is omitted here in favor
// of being embedded directly into the templates. (The charset directive is a
// good example. It is and should be set to utf-8, always, so there's no reason
// to store and process it on a per-page basis.)
type Meta struct {
	ID           uint32
	Title        string
	Keywords     string
	Description  string
	Language     string
	PageTopic    string
	RevisitAfter string
	Robots       string
}

// New creates a new Page object with usable defaults.
func New() *Page {
	return &Page{Meta: &Meta{}, Template: tmpl.New()}
}
