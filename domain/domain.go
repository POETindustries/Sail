package domain

import (
	"fmt"
	"sail/tmpl"
)

// Domain is the primary type for determining the presentation
// of page data. It holds references to page metadata for use
// in an html document's head area and to the template that
// controls the actual layout.
type Domain struct {
	ID       uint32
	Name     string
	Meta     *Meta
	Template *tmpl.Template
}

// String prints the domain's data in an easily readable format.
func (d *Domain) String() string {
	str := "DOMAIN '%s': {ID:%d | Meta:%+v | Template:%+v}"
	return fmt.Sprintf(str, d.Name, d.ID, d.Meta, d.Template)
}

// New creates a basic Domain object filled with default values.
func New() *Domain {
	return &Domain{Meta: &Meta{}, Template: tmpl.New()}
}
