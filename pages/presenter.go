package pages

import (
	"bytes"
	"fmt"
	"html/template"
	"sail/cache"
	"sail/domains"
	"sail/page"
	"sail/widget"
	"strings"
)

// Presenter initiates page creation and loading for handling requests
// by users from the www.
type Presenter struct {
	page   *page.Page
	markup *bytes.Buffer
	url    string
}

func (p *Presenter) compile() error {
	return p.page.Domain.Template.Execute(p.markup, p)
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
func (p *Presenter) MetaTitle() string { return p.page.Domain.Meta.Title }

// MetaKeywords returns the keywords for use in the html <meta> tag
// within the <head> area.
func (p *Presenter) MetaKeywords() string { return p.page.Domain.Meta.Keywords }

// MetaDescription returns the page's description for use in the html
// <meta> tag within the <head> area.
func (p *Presenter) MetaDescription() string { return p.page.Domain.Meta.Description }

// MetaLanguage returns the language value for use in the html <meta>
// tag within the <head> area.
func (p *Presenter) MetaLanguage() string { return p.page.Domain.Meta.Language }

// MetaPageTopic returns the page-topic value for use in the html
// <meta> tag within the <head> area.
func (p *Presenter) MetaPageTopic() string { return p.page.Domain.Meta.PageTopic }

// MetaRevisit returns the desired crawler revisit value for use in
//the html <meta> tag within the <head> area.
func (p *Presenter) MetaRevisit() string { return p.page.Domain.Meta.RevisitAfter }

// MetaRobots returns the desired robots value for use in the html
// <meta> tag within the <head> area.
func (p *Presenter) MetaRobots() string { return p.page.Domain.Meta.Robots }

// Widget returns a pointer to the widget designated by the name
// parameter. If no such widget exists, an empty widget is returned.
func (p *Presenter) Widget(name string) (w *widget.Widget) {
	if w = p.page.Domain.Template.Widgets[name]; w == nil {
		return widget.New()
	}
	return
}

// Menu returns the menu identified by the name, if possible.
// It is guaranteed to return an object of the correct type; if the
// desired object does not exist, an empty object is returned with
// all necessary components minimally initialized.
func (p *Presenter) Menu(name string, isMain bool) *widget.Menu {
	w := p.Widget(name)
	m, ok := w.Data.(*widget.Menu)
	if !ok {
		return &widget.Menu{Entries: []*widget.MenuEntry{}}
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

// New creates a new presenter object with all necessary fields properly
// initialized.
func New() *Presenter {
	return &Presenter{page: page.New(), markup: bytes.NewBufferString("")}
}

func NewFromCache(url string) *Presenter {
	page, ok := cache.Pages[url].(*page.Page)
	if ok {
		fmt.Printf("page found in cache: %d\n", page.ID)
		return &Presenter{page: page, markup: bytes.NewBufferString("")}
	}
	return NewWithURL(url)
}

// NewWithURL expects a valid request uri in order to compile the
// corresponding page data. It is guaranteed to retun a functioning
// presenter object even if the url parameter does not lead to any data.
func NewWithURL(url string) *Presenter {
	if len(url) <= 1 {
		return NewWithID(1)
	}
	presenter := New()
	presenter.url = url
	pages, err := fetchByURL(url)
	if len(pages) == 0 || err != nil {
		return NewWithID(1)
	}
	pages[0].Domain = domains.FromCache(pages[0].Domain.ID)
	presenter.page = pages[0]
	cache.Pages[url] = pages[0]
	fmt.Printf("page added to cache: %d\n", pages[0].ID)
	return presenter
}

// NewWithID expects an id value in order to compile the
// corresponding page data. It is guaranteed to retun a functioning
// presenter object even if the id parameter does not lead to any data.
func NewWithID(id uint32) *Presenter {
	presenter := New()
	pages, err := fetchByID(id)
	if len(pages) == 0 || err != nil {
		presenter.page = load404()
	} else {
		pages[0].Domain = domains.BuildWithID(pages[0].Domain.ID)[0]
		presenter.page = pages[0]
	}
	return presenter
}
