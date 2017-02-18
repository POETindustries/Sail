package group

import (
	"net/http"
	"net/url"
	"sail/session"
	"sail/user"
	"sail/user/rights"
)

// Bouncer enforces access to resources and actions. It checks if
// users have access to the url they request and if they are allowed
// to perform the actions requested in either GET or POST queries.
//
// In order to check for these permissions, the bouncer needs a
// user's id, which it can find out by its own from a given request
// or session, or it can be passed the id directly.
type Bouncer struct {
	req *http.Request
}

// NewBouncer returns a Bouncer initialized with the given request.
func NewBouncer(req *http.Request) *Bouncer {
	return &Bouncer{req: req}
}

// Pass checks for access violations with the given user id.
func (b *Bouncer) Pass(uid uint32) bool {
	d, err := rights.Dom(b.req.URL.Path)
	if err != nil {
		return true
	}
	p := All().Permission(uid, d)
	return b.validateGET(p) && b.validatePOST(p)
}

// PassByCookie checks for access violations with the information
// gained from the given cookie.
func (b *Bouncer) PassByCookie(c *http.Cookie) bool {
	s := session.DB().Get(c.Value)
	if s == nil {
		return false
	}
	return b.PassBySession(session.DB().Get(c.Value))
}

// PassBySession expects a session and checks if the user designated
// by the session has the right to access the path and execute the
// actions specified in the Bouncer's request object.
func (b *Bouncer) PassBySession(s *session.Session) bool {
	return b.Pass(user.LoadNew(s.User).ID())
}

// Sanitize cleans the Bouncer's request object of unwanted and
// invalid content. It expects the path of a content object that
// should be requested instead of the offending url.
func (b *Bouncer) Sanitize(path string) {
	b.req.URL.Path = path
	b.req.URL.RawQuery = ""
	b.req.PostForm = nil
	b.req.Form = nil
}

func (b *Bouncer) validateGET(p *rights.Permission) bool {
	return b.validate(b.req.URL.Query(), p)
}

func (b *Bouncer) validatePOST(p *rights.Permission) bool {
	return b.validate(b.req.PostForm, p)
}

func (b *Bouncer) validate(vals url.Values, p *rights.Permission) bool {
	if !p.R() {
		return false
	}
	if !p.C() {
		// TODO: implement test for create-operations
	}
	if !p.U() {
		// TODO: implement test for update-operations
	}
	if !p.D() {
		// TODO: implement test for delete-operations
	}
	return true
}
