/******************************************************************************
Copyright 2015-2017 POET Industries

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the
following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
******************************************************************************/

package session

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
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

// Database manages an internal list of active sessions.
type Database struct {
	sessions map[string]*Session
	seeds    []byte
	quota    int
}

var db *Database

// DB returns a pointer to the session database singleton.
func DB() *Database {
	if db == nil {
		db = &Database{sessions: make(map[string]*Session)}
		db.reseed()
	}
	return db
}

// Add adds a new session to the session pool.
func (db *Database) Add(session *Session) {
	db.sessions[session.ID] = session
}

// Clean removes all expired sessions.
func (db *Database) Clean() {
	// TODO: expiration time is hardcoded to 6 hours.
	// This needs to be a config setting in the future.
	for _, s := range db.sessions {
		if time.Since(s.Time).Hours() > 6 {
			db.Remove(s.ID)
		}
	}
}

// Get returns the session specified by the given id.
func (db *Database) Get(id string) *Session {
	return db.sessions[id]
}

// Has returns true if the session specified by the given id exists
// in the database of active sessions.
func (db *Database) Has(id string) bool {
	return db.sessions[id] != nil
}

// ID ...is here for whatever reason.
func (db *Database) ID(id string) (i string) {
	if s := db.sessions[id]; s != nil {
		i = s.ID
	}
	return
}

// Lang returns the session's language.
func (db *Database) Lang(id string) (l string) {
	if s := db.sessions[id]; s != nil {
		l = s.Lang
	}
	return
}

// Remove removes the session with the given id from the session pool.
func (db *Database) Remove(id string) {
	delete(db.sessions, id)
}

// Seed returns the next available seed for use in creating session ids.
func (db *Database) Seed() (b []byte) {
	if len(db.seeds) < 5 && !db.reseed() {
		return // TODO fill b with random data from a secondary source?
	}
	b = db.seeds[:4]
	DB().seeds = DB().seeds[5:]
	return b
}

// Start resets a session's timer to the current time.
func (db *Database) Start(id string) {
	if s := db.sessions[id]; s != nil {
		s.Start()
	}
}

// User returns the user name associated with the session specified
// by the given id.
func (db *Database) User(id string) (username string) {
	if s := db.sessions[id]; s != nil {
		username = s.User
	}
	return
}

func (db *Database) reseed() bool {
	db.checkSeedQuota()
	if db.quota < 10000 { // are there enough bytes left to download another chunk?
		return false
	}
	resp, err := http.Get("https://www.random.org/strings/?num=4&len=4&digits=on&upperalpha=on&loweralpha=on&format=plain") // TODO num needs to be higher in production, maybe make this editable through config file?
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return false
	}
	db.seeds, _ = ioutil.ReadAll(resp.Body)
	return true
}

func (db *Database) checkSeedQuota() bool {
	resp, err := http.Get("https://www.random.org/quota/?format=plain")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if db.quota, err = strconv.Atoi(string(body[:len(body)-1])); err != nil {
		db.quota = 0
	}
	return true
}
