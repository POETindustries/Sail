package session

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User describes the base interface that concrete structs need
// to satisfy in order to work with Sail's session framework.
type User interface {
	ID() uint32
	Name() string
	Load() bool
	Hash() (string, error)
}

// UserDB is a simple in-memory database of all users that are
// currently active. It helps to improve performance and reduce
// queries to persistent storage. Operations on this object
// do not change persistent user data.
type UserDB struct {
	names map[string]User
	ids   map[uint32]User
}

var userdb *UserDB

// Users returns a pointer to the system-wide user database.
func Users() *UserDB {
	if userdb == nil {
		userdb = &UserDB{
			names: map[string]User{},
			ids:   map[uint32]User{}}
	}
	return userdb
}

// Add inserts a user into the database. Previous user objects
// with the same name or id are overwritten.
func (db *UserDB) Add(u User) {
	db.names[u.Name()] = u
	db.ids[u.ID()] = u
}

// Has checks if a user is already in the database.
func (db *UserDB) Has(u User) bool {
	return db.HasID(u.ID()) || db.HasName(u.Name())
}

// HasName checks if a user with the given name exists in the
// database.
func (db *UserDB) HasName(name string) bool {
	_, ok := db.names[name]
	return ok
}

// HasID checks if a user with the given id exists in the
// database.
func (db *UserDB) HasID(id uint32) bool {
	_, ok := db.ids[id]
	return ok
}

// ByName fetches a user from the database that matches the
// username given. returns nil if none was found.
func (db *UserDB) ByName(name string) User {
	return db.names[name]
}

// ByID fetches a user from the database that matches the
// id given. returns nil if none was found.
func (db *UserDB) ByID(id uint32) User {
	return db.ids[id]
}

// All returns all users the system knows about.
func (db *UserDB) All() []User {
	var us []User
	for _, v := range db.ids {
		us = append(us, v)
	}
	return us
}

// Remove deletes the user from the database.
func (db *UserDB) Remove(u User) {
	db.RemoveID(u.ID())
}

// RemoveName deletes the user that matches the given name
// from the database.
func (db *UserDB) RemoveName(name string) {
	if db.HasName(name) {
		id := db.ByName(name).ID()
		delete(db.names, name)
		delete(db.ids, id)
	}
}

// RemoveID deletes the user that matches the given id from
// the database.
func (db *UserDB) RemoveID(id uint32) {
	if db.HasID(id) {
		name := db.ByID(id).Name()
		delete(db.names, name)
		delete(db.ids, id)
	}
}

// Verify returns a User and a boolean value indicating whether user
// and password match entries in the user database. If so, the User
// return value is safe for use.
//
// Since there has to be a database request anyway, the resulting
// User object is returned in order to give the opportunity to use it
// right away instead of fetching anew a few lines later.
func Verify(u User, pass string) (User, bool) {
	if pass == "" || !u.Load() {
		return nil, false
	}
	fmt.Printf("%+v\n", u)
	if hash, err := u.Hash(); err == nil {
		if bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil {
			return u, true
		}
	}
	return nil, false
}

// Encrypt generates a cryptographically secure hash from the
// given string. It is strong enough for password storage.
func Encrypt(s string) string {
	if k, err := bcrypt.GenerateFromPassword([]byte(s), 12); err == nil {
		return fmt.Sprintf("%s\n", k)
	}
	return ""
}
