package page

import (
	"html/template"
)

type Meta struct {
	title       string
	keywords    string
	description string
	robots      string
}

type Page struct {
	meta    *Meta
	frame   *template.Template
	content *template.Template
}
