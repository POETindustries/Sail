package backend

import (
	"bytes"
	"sail/conf"
	"sail/errors"
	"sail/page/fallback"
	"sail/page/template"
	"sail/user"
	"sail/user/session"
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
	Session *session.Session
	User    *user.User

	msg      string
	url      string
	template *template.Template
}

// New creates a new presenter object with all necessary
// fields properly initialized.
func New(s *session.Session) *Presenter {
	p := &Presenter{Session: s, template: template.New()}
	p.template.Name = "default-backend"
	if s != nil {
		p.User = user.ByName(s.User)
	}
	return p
}

func (p *Presenter) Compile() *bytes.Buffer {
	var markup bytes.Buffer
	if err := p.template.Execute(&markup, p); err != nil {
		errors.Log(err, conf.Instance().DevMode)
		return bytes.NewBufferString(fallback.NOTFOUND404)
	}
	return &markup
}

func (p *Presenter) Message() string {
	return p.msg
}

func (p *Presenter) SetMessage(msg string) {
	p.msg = msg
}

func (p *Presenter) URL() string {
	if p.url == "/office/login" && p.Session != nil {
		return "/office/"
	}
	return p.url
}

func (p *Presenter) SetURL(url string) {
	p.url = url
}

/*
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
*/
