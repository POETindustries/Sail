package session

import "time"

// Database manages an internal list of active sessions.
type Database struct {
	sessions map[string]*Session
}

var instance *Database

func new() *Database {
	return &Database{sessions: make(map[string]*Session)}
}

// DB returns a pointer to the session database singleton.
func DB() *Database {
	if instance == nil {
		instance = new()
	}
	return instance
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

// Start resets a session's timer to the current time.
func (db *Database) Start(id string) {
	if s := db.sessions[id]; s != nil {
		s.Start()
	}
}

// User returns the user name associated with the session specified
// by the given id.
func (db *Database) User(id string) (u string) {
	if s := db.sessions[id]; s != nil {
		u = s.User
	}
	return
}
