package session

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Session represents a user session and contains information for
// correctly displaying localization information and other session-
// related data.
type Session struct {
	ID   string
	User string
	Lang string
	Time time.Time
}

// New returns a new Session and sets its internal state according
// to the given request and user data.
func New(req *http.Request, user string) (s *Session) {
	s = &Session{User: user}
	s.setLang(req.Header.Get("Accept-Language"))
	s.genID()
	return s
}

// Start resets the Session's timer to the current time.
func (s *Session) Start() {
	s.Time = time.Now()
}

func (s *Session) setLang(lang string) {
	s.Lang = strings.Split(lang, ",")[0]
}

func (s *Session) genID() {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		panic(err.Error())
	}
	b = append(b, DB().Seed()...)
	s.ID = fmt.Sprintf("%x", sha1.Sum(b))
}
