package session

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"net/http"
	"strconv"
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
	if s.genID(req.Header.Get("X-Forwarded-For")) {
		return s
	}
	return
}

// Start resets the Session's timer to the current time.
func (s *Session) Start() {
	s.Time = time.Now()
}

func (s *Session) setLang(lang string) {
	s.Lang = strings.Split(lang, ",")[0]
}

func (s *Session) genID(ip string) bool {
	b := make([]byte, 10)
	if _, err := rand.Read(b); err == nil {
		b = append(b, []byte(ip)...)
		b = append(b, []byte(strconv.Itoa(time.Now().Nanosecond()))...)
		s.ID = fmt.Sprintf("%x", sha1.Sum(b))

		return true
	}
	return true
}
