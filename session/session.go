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

type Session struct {
	ID   string
	User string
	Lang string
	Time time.Time
}

func New(req *http.Request, user string) (s *Session) {
	s = &Session{User: user}
	s.setLang(req.Header.Get("Accept-Language"))
	if s.genID(req.Header.Get("X-Forwarded-For")) {
		return s
	}
	return
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
