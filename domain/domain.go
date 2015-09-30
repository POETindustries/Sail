package domain

import (
	"sail/storage"
	"sail/tmpl"
)

type Domain struct {
	ID       uint32
	Name     string
	Meta     *Meta
	Template *tmpl.Template
}

func (d *Domain) ScanFromDB() bool {
	m := Meta{}
	data := storage.Instance().DomainData(d.ID)

	if data.Scan(&d.ID, &d.Name, &m.Title, &m.Keywords, &m.Description, &m.Language,
		&m.PageTopic, &m.RevisitAfter, &m.Robots, &d.Template) != nil {
		return false
	}
	d.Meta = &m

	return true
}

func (d *Domain) loadTemplate() {
	d.Template = tmpl.New("")
}
