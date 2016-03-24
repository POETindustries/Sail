package session

import "time"

// Database manages an internal list of active sessions.
type Database struct {
	sessions map[string]*Session
}

var instance *Database

func new() *Database {
	db := Database{}
	db.sessions = make(map[string]*Session)
	return &db
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

// Remove removes the session with the given id from the session pool.
func (db *Database) Remove(id string) {
	delete(db.sessions, id)
}

// ID ...is here for whatever reason.
func (db *Database) ID(id string) string {
	return db.sessions[id].ID
}

// User returns the user name associated with the session specified
// by the given id.
func (db *Database) User(id string) string {
	return db.sessions[id].User
}

// Get returns the session specified by the given id.
func (db *Database) Get(id string) *Session {
	return db.sessions[id]
}

// Lang returns the session's language.
func (db *Database) Lang(id string) string {
	return db.sessions[id].Lang
}

// Has returns true if the session specified by the given id exists
// in the database of active sessions.
func (db *Database) Has(id string) bool {
	return db.sessions[id] != nil
}

// Start resets a session's timer to the current time.
func (db *Database) Start(id string) {
	db.sessions[id].Start()
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
