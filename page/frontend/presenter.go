package frontend

import (
	"bytes"
	"html/template"
	"regexp"
	"sail/conf"
	"sail/errors"
	"sail/file"
	"sail/page/content"
	"sail/page/fallback"
	tpl "sail/page/template"
	"sail/page/widget"
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
	msg      string
	url      string
	content  *content.Content
	template *tpl.Template
}

// New creates a new presenter object with all necessary
// fields properly initialized.
func New(cnt *content.Content, tmpl *tpl.Template) *Presenter {
	return &Presenter{
		content:  cnt,
		template: tmpl}
}

func (p *Presenter) Compile() *bytes.Buffer {
	var markup bytes.Buffer
	if err := p.template.Execute(&markup, p); err != nil {
		errors.Log(err, conf.Instance().DevMode)
		return bytes.NewBufferString(fallback.NOTFOUND404)
	}
	b := markup.Bytes()
	p.replaceInternalLinks(&b)
	return bytes.NewBuffer(b)
}

func (p *Presenter) Message() string {
	return p.msg
}

func (p *Presenter) SetMessage(msg string) {
	p.msg = msg
}

func (p *Presenter) URL() string {
	return p.url
}

func (p *Presenter) SetURL(url string) {
	p.url = url
}

// PageTitle returns the title of the currently held page object.
func (p *Presenter) PageTitle() string { return p.content.Title }

// PageOwner returns the name of the page's owner.
func (p *Presenter) PageOwner() string { return p.content.Owner }

// PageEditDate returns a format string for the date the page was
// edited last.
func (p *Presenter) PageEditDate() string { return p.content.EDate }

// PageCreateDate returns a string-formatted representation of the
// date the page was created.
func (p *Presenter) PageCreateDate() string { return p.content.CDate }

// PageContent returns the page's contents in an html-encoded format.
func (p *Presenter) PageContent() template.HTML {
	return template.HTML(p.content.Content)
}

// MetaTitle returns the page title for use in the html <meta> tag
// within the <head> area.
func (p *Presenter) MetaTitle() string { return p.content.Meta.Title }

// MetaKeywords returns the keywords for use in the html <meta> tag
// within the <head> area.
func (p *Presenter) MetaKeywords() string { return p.content.Meta.Keywords }

// MetaDescription returns the page's description for use in the html
// <meta> tag within the <head> area.
func (p *Presenter) MetaDescription() string { return p.content.Meta.Description }

// MetaLanguage returns the language value for use in the html <meta>
// tag within the <head> area.
func (p *Presenter) MetaLanguage() string { return p.content.Meta.Language }

// MetaPageTopic returns the page-topic value for use in the html
// <meta> tag within the <head> area.
func (p *Presenter) MetaPageTopic() string { return p.content.Meta.PageTopic }

// MetaRevisit returns the desired crawler revisit value for use in
//the html <meta> tag within the <head> area.
func (p *Presenter) MetaRevisit() string { return p.content.Meta.RevisitAfter }

// MetaRobots returns the desired robots value for use in the html
// <meta> tag within the <head> area.
func (p *Presenter) MetaRobots() string { return p.content.Meta.Robots }

// Widget returns a pointer to the widget designated by the name
// parameter. If no such widget exists, an empty widget is returned.
func (p *Presenter) Widget(name string) (w *widget.Widget) {
	if w = p.template.Widgets[name]; w == nil {
		return widget.New()
	}
	return
}

// Menu returns the menu identified by the name, if possible.
// It is guaranteed to return an object of the correct type; if the
// desired object does not exist, an empty object is returned with
// all necessary components minimally initialized.
func (p *Presenter) NavMenu(name string, isMain bool) *widget.Nav {
	w := p.Widget(name)
	m, ok := w.Data.(*widget.Nav)
	if !ok {
		return &widget.Nav{Entries: []*widget.NavEntry{}}
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
	t, ok := w.Data.(*widget.Text)
	if ok {
		return template.HTML(t.Content)
	}
	return template.HTML("")
}

func (p *Presenter) replaceInternalLinks(markup *[]byte) {
	r, _ := regexp.Compile("=\"uuid/[0-9]+\"")
	refs := make(map[string]bool)
	for _, r := range r.FindAll(*markup, -1) {
		refs[string(r[2:len(r)-1])] = true
	}
	for k := range refs {
		*markup = bytes.Replace(*markup, []byte(k), file.StaticAddr(k), -1)
	}
}
