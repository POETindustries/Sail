package backend

import (
	"bytes"
	"sail/conf"
	"sail/errors"
	"sail/file"
	"sail/object/fallback"
	"sail/object/template"
	"sail/object/widget"
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
	Session     *session.Session
	User        *user.User
	FileManager *file.Manager

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

// Compile creates html markup from the prsenter's template
// and the data stored in the presenter at the time of
// compilation. The resulting markup is ready to be sent
// as response to the last http request.
func (p *Presenter) Compile() *bytes.Buffer {
	var markup bytes.Buffer
	if err := p.template.Execute(&markup, p); err != nil {
		errors.Log(err, conf.Instance().DevMode)
		return bytes.NewBufferString(fallback.NOTFOUND404)
	}
	return &markup
}

// Message returns the currently saved message or an empty
// string if no message is set.
func (p *Presenter) Message() string {
	return p.msg
}

// SetMessage allows passing a message from processing the
// request to the presenter for displaying in the web page.
func (p *Presenter) SetMessage(msg string) {
	p.msg = msg
}

// URL returns the url currently associated with the presenter.
func (p *Presenter) URL() string {
	if p.url == "/office/login" && p.Session != nil {
		return "/office/"
	}
	return p.url
}

// SetURL should be used to change the presenter's internal
// url after it has already been initialized.
func (p *Presenter) SetURL(url string) {
	if p.url != url {
		p.url = url
		if p.url == "/office/content" {
			p.FileManager = file.NewManager("/")
		}
	}
}

// MainMenu returns available menu entry data, depending on
// the user's access permissions.
func (p *Presenter) MainMenu() *widget.Nav {
	for _, e := range p.mainMenu.Entries {
		e.Active = p.url == e.RefURL
	}
	return p.mainMenu
}

func (p *Presenter) buildNav(uid uint32) *widget.Nav {
	nav := widget.Nav{Entries: []*widget.NavEntry{&widget.NavEntry{
		ID: 1, Name: "Home", RefURL: "/office/"}}}
	if group.All().Mode(uid, rights.Content) > 1 {
		nav.Entries = append(nav.Entries, &widget.NavEntry{
			ID: 2, Name: "Content", RefURL: "/office/content"})
	}
	if group.All().Mode(uid, rights.Users) > 1 {
		nav.Entries = append(nav.Entries, &widget.NavEntry{
			ID: 3, Name: "Users & Groups", RefURL: "/office/users"})
	}
	if group.All().Mode(uid, rights.Config) > 1 {
		nav.Entries = append(nav.Entries, &widget.NavEntry{
			ID: 4, Name: "Configuration", RefURL: "/office/config"})
	}
	if group.All().Mode(uid, rights.Maintenance) > 1 {
		nav.Entries = append(nav.Entries, &widget.NavEntry{
			ID: 5, Name: "Maintenance", RefURL: "/office/maintenance"})
	}
	return &nav
}
