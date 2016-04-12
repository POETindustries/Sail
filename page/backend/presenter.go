package backend

import (
	"bytes"
	"sail/conf"
	"sail/errors"
	"sail/page/fallback"
	"sail/page/template"
	"sail/page/widget"
	"sail/user"
	"sail/user/group"
	"sail/user/rights"
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

	msg string
	url string

	template *template.Template
	mainMenu *widget.Nav
}

// New creates a new presenter object with all necessary
// fields properly initialized.
func New(s *session.Session, u *user.User) *Presenter {
	p := &Presenter{Session: s, User: u, template: template.New()}
	p.template.Name = "default-backend"
	if u != nil {
		p.mainMenu = p.buildNav(p.User.ID)
	} else if s != nil {
		p.User = user.ByName(p.Session.User)
		p.mainMenu = p.buildNav(p.User.ID)
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

func (p *Presenter) MainMenu() *widget.Nav {
	for _, e := range p.mainMenu.Entries {
		e.Active = p.url == e.RefURL
	}
	return p.mainMenu
}

func (p *Presenter) buildNav(uid uint32) *widget.Nav {
	home := &widget.NavEntry{
		ID: 1, Name: "Home", RefURL: "/office/"}
	content := &widget.NavEntry{
		ID: 2, Name: "Content", RefURL: "/office/content"}
	user := &widget.NavEntry{
		ID: 3, Name: "Users & Groups", RefURL: "/office/users"}
	config := &widget.NavEntry{
		ID: 4, Name: "Configuration", RefURL: "/office/config"}
	maintenance := &widget.NavEntry{
		ID: 5, Name: "Maintenance", RefURL: "/office/maintenance"}

	nav := widget.Nav{Entries: []*widget.NavEntry{home}}
	if group.All().Mode(uid, rights.Content) > 1 {
		nav.Entries = append(nav.Entries, content)
	}
	if group.All().Mode(uid, rights.Users) > 1 {
		nav.Entries = append(nav.Entries, user)
	}
	if group.All().Mode(uid, rights.Config) > 1 {
		nav.Entries = append(nav.Entries, config)
	}
	if group.All().Mode(uid, rights.Maintenance) > 1 {
		nav.Entries = append(nav.Entries, maintenance)
	}
	return &nav
}
