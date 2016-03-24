package session

import "time"

type Database struct {
	sessions map[string]*Session
}

var instance *Database

func new() *Database {
	db := Database{}
	db.sessions = make(map[string]*Session)
	return &db
}

func DB() *Database {
	if instance == nil {
		instance = new()
	}
	return instance
}

func (db *Database) Add(session *Session) {
	db.sessions[session.ID] = session
}

func (db *Database) Remove(id string) {
	delete(db.sessions, id)
}

func (db *Database) ID(id string) string {
	return db.sessions[id].ID
}

func (db *Database) User(id string) string {
	return db.sessions[id].User
}

func (db *Database) Get(id string) *Session {
	return db.sessions[id]
}

func (db *Database) Lang(id string) string {
	return db.sessions[id].Lang
}

func (db *Database) Has(id string) bool {
	return db.sessions[id] != nil
}

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
