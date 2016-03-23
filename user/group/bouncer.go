package group

import (
	"net/http"
	"sail/user"
	"sail/user/rights"
	"sail/user/session"
)

// Bouncer enforces access to resources and actions.
type Bouncer struct {
	req *http.Request
}

// NewBouncer returns a Bouncer initialized with the given request.
func NewBouncer(req *http.Request) *Bouncer {
	return &Bouncer{req: req}
}

// Pass checks for access violations with nothing more than the
// Bouncer's request object.
func (b *Bouncer) Pass() bool {
	c, _ := b.req.Cookie("session")
	if c == nil {
		return true
	}
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
	return b.PassByUser(u.ID)
}

func (b *Bouncer) PassByUser(uid uint32) bool {
	d, err := rights.Dom(b.req.URL.Path)
	return err != nil || Cache().Permission(uid, d).R()
}
