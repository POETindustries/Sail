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
	"sail/conf"
	"strconv"
	"strings"
	"sync"
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
	return s
}

// Start resets the Session's timer to the current time.
func (s *Session) Start() {
	s.Time = time.Now()
}

func (s *Session) setLang(lang string) {
	s.Lang = strings.Split(lang, ",")[0]
}

// Database manages an internal list of active sessions.
type Database struct {
	sync.RWMutex
	sessions map[string]*Session
	seeds    []byte
	quota    int
}

var db *Database
var sessionInit sync.Once

// DB returns a pointer to the session database singleton.
func DB() *Database {
	sessionInit.Do(func() {
		db = &Database{sessions: make(map[string]*Session)}
		db.reseed()
	})
	return db
}

// Clean removes all expired sessions.
func (db *Database) Clean() {
	// TODO: expiration time is hardcoded to 6 hours.
	// This needs to be a config setting in the future.
	db.Lock()
	for _, s := range db.sessions {
		if time.Since(s.Time).Hours() > 6 {
			Users().RemoveName(s.User)
			delete(db.sessions, s.ID)
		}
	}
	db.Unlock()
}

// Has returns true if the session specified by the given id exists
// in the database of active sessions.
func (db *Database) Has(id string) bool {
	db.RLock()
	defer db.RUnlock()
	return db.sessions[id] != nil
}

// Lang returns the session's language.
func (db *Database) Lang(id string) (l string) {
	db.RLock()
	if s := db.sessions[id]; s != nil {
		l = s.Lang
	}
	db.RUnlock()
	return
}

// New creates a new session, stores it in the database and
// returns the unique session id.
func (db *Database) New(req *http.Request, user string) *http.Cookie {
	s := New(req, user)
	s.Start()
	db.Lock()
	s.ID = db.nextID()
	db.sessions[s.ID] = s
	c := &http.Cookie{Name: "id", Value: s.ID}
	db.Unlock()
	if !conf.Instance().DevMode {
		c.Secure = true
		c.HttpOnly = true
	}
	return c
}

// Remove removes the session with the given id from the session pool.
func (db *Database) Remove(id string) {
	db.Lock()
	// TODO 2017-02-26: Danger Zone: Remove can deadlock
	// if the user database is locked at this time and
	// wants to act on the session database, which will
	// be locked by Remove. This is never the case in
	// the current implementation and it should never be.
	// The user database has no business being on
	// session database's lawn. Still, something worth
	// looking out for. The same applies to Clean().
	Users().RemoveName(db.sessions[id].User)
	delete(db.sessions, id)
	db.Unlock()
}

// Start resets a session's timer to the current time.
func (db *Database) Start(id string) {
	db.Lock()
	if s := db.sessions[id]; s != nil {
		s.Start()
	}
	db.Unlock()
}

// User returns the user name associated with the session specified
// by the given id.
func (db *Database) User(id string) (username string) {
	db.RLock()
	if s := db.sessions[id]; s != nil {
		username = s.User
	}
	db.RUnlock()
	return
}

func (db *Database) nextID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		panic(err.Error())
	}
	if len(db.seeds) > 5 || db.reseed() {
		b = append(b, db.seeds[:4]...)
		db.seeds = db.seeds[5:]
	}
	return fmt.Sprintf("%x", sha1.Sum(b))
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
