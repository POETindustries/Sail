package response

import (
	"bytes"
	"html/template"
	"sail/page/data"
	"strings"
)

// Presenter initiates page creation and loading for handling requests
// by users from the www. It also serves as the content provider for
// templates.
//
// All exported functions and fields that return strings and
// string-derived types are safe for use inside a template. All exported
// functions and fields of type bool are safe for use as conditions
// inside templates.
type Presenter struct {
	HasMessage bool
	Message    string
	FallbackID uint32

	page   *data.Page
	markup *bytes.Buffer
	url    string
}

// New creates a new presenter object with all necessary fields properly
// initialized.
func NewPresenter() *Presenter {
	return &Presenter{
		FallbackID: 1,
		page:       data.NewPage(),
		markup:     bytes.NewBufferString("")}
}

func (p *Presenter) Compile() (*bytes.Buffer, error) {
	err := p.page.Template.Execute(p.markup, p)
	return p.markup, err
}

// PageTitle returns the title of the currently held page object.
func (p *Presenter) PageTitle() string { return p.page.Title }

// PageOwner returns the name of the page's owner.
func (p *Presenter) PageOwner() string { return p.page.Owner }

// PageEditDate returns a format string for the date the page was
// edited last.
func (p *Presenter) PageEditDate() string { return p.page.EDate }

// PageCreateDate returns a string-formatted representation of the
// date the page was created.
func (p *Presenter) PageCreateDate() string { return p.page.CDate }

// PageContent returns the page's contents in an html-encoded format.
func (p *Presenter) PageContent() template.HTML {
	return template.HTML(p.page.Content)
}

// MetaTitle returns the page title for use in the html <meta> tag
// within the <head> area.
func (p *Presenter) MetaTitle() string { return p.page.Meta.Title }

// MetaKeywords returns the keywords for use in the html <meta> tag
// within the <head> area.
func (p *Presenter) MetaKeywords() string { return p.page.Meta.Keywords }

// MetaDescription returns the page's description for use in the html
// <meta> tag within the <head> area.
func (p *Presenter) MetaDescription() string { return p.page.Meta.Description }

// MetaLanguage returns the language value for use in the html <meta>
// tag within the <head> area.
func (p *Presenter) MetaLanguage() string { return p.page.Meta.Language }

// MetaPageTopic returns the page-topic value for use in the html
// <meta> tag within the <head> area.
func (p *Presenter) MetaPageTopic() string { return p.page.Meta.PageTopic }

// MetaRevisit returns the desired crawler revisit value for use in
//the html <meta> tag within the <head> area.
func (p *Presenter) MetaRevisit() string { return p.page.Meta.RevisitAfter }

// MetaRobots returns the desired robots value for use in the html
// <meta> tag within the <head> area.
func (p *Presenter) MetaRobots() string { return p.page.Meta.Robots }

// Widget returns a pointer to the widget designated by the name
// parameter. If no such widget exists, an empty widget is returned.
func (p *Presenter) Widget(name string) (w *data.Widget) {
	if w = p.page.Template.Widgets[name]; w == nil {
		return data.NewWidget()
	}
	return
}

// Menu returns the menu identified by the name, if possible.
// It is guaranteed to return an object of the correct type; if the
// desired object does not exist, an empty object is returned with
// all necessary components minimally initialized.
func (p *Presenter) Menu(name string, isMain bool) *data.Menu {
	w := p.Widget(name)
	m, ok := w.Data.(*data.Menu)
	if !ok {
		return &data.Menu{Entries: []*data.MenuEntry{}}
	}
	if isMain {
		for _, e := range m.Entries {
			e.Active = strings.HasPrefix(p.url, e.RefURL)
		}
	}
	return m
}

// TextWidget returns the text of the text widget identified by the
// name parameter. It is guaranteed to return an object of the correct
// type; if the desired object doesn't exist, returns an empty string.
func (p *Presenter) TextWidget(name string) template.HTML {
	w := p.Widget(name)
	t, ok := w.Data.(*data.Text)
	if ok {
		return template.HTML(t.Content)
	}
	return template.HTML("")
}
