package group

import (
	"net/http"
	"sail/user"
	"sail/user/rights"
	"sail/user/session"
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
	return err != nil || Cache().Permission(uid, d).R()
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
	u := user.ByName(s.User)
	if u == nil {
		return false
	}
	return b.Pass(u.ID)
}

func (b *Bouncer) Sanitize(path, query string) {
	b.req.URL.Path = path
	b.req.URL.RawQuery = query
	b.req.PostForm = nil
}
