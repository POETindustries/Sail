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
	"fmt"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// User describes the base interface that concrete structs need
// to satisfy in order to work with Sail's session framework.
type User interface {
	Copy() User
	ID() uint32
	Name() string
}

// UserDB is a simple in-memory database of all users that are
// currently active. It helps to improve performance and reduce
// queries to persistent storage. Operations on this object
// do not change persistent user data.
//
// Concurrency Safety
//
// UserDB itself is safe for concurrent use from multiple
// goroutines. In order to preserve safety, when methods return
// an object directly from the database, they actually return
// a copy of the object. This is enforced by interface User's
// Copy() method.
//
// Writes to objects retrieved this way do therefore not
// propagate to the objects inside the database. If such an
// object is mutated, the related object in the database has
// to be overwritten with the newly changed one.
type UserDB struct {
	sync.RWMutex
	names map[string]User
	ids   map[uint32]User
}

var userdb *UserDB
var userdbInit sync.Once

// NewUserDB creates and returns an empty user cache.
func NewUserDB() *UserDB {
	return &UserDB{
		names: map[string]User{},
		ids:   map[uint32]User{}}
}

// Add inserts a user into the database. Previous user objects
// with the same name or id are overwritten. Add is suitable
// for updating an entry in the database.
func (db *UserDB) Add(u User) {
	new := u.Copy()
	db.Lock()
	db.names[strings.ToLower(u.Name())] = new
	db.ids[u.ID()] = new
	db.Unlock()
}

// All returns the names of all users that are currently cached.
func (db *UserDB) All() []string {
	var us []string
	db.RLock()
	for _, v := range db.ids {
		us = append(us, v.Name())
	}
	db.RUnlock()
	return us
}

// Has checks if a user is already in the database.
func (db *UserDB) Has(u User) bool {
	db.RLock()
	defer db.RUnlock()
	return db.hasID(u.ID()) || db.hasName(u.Name())
}

// HasName checks if a user with the given name exists in the
// database.
func (db *UserDB) HasName(name string) bool {
	db.RLock()
	defer db.RUnlock()
	return db.hasName(name)
}

func (db *UserDB) hasName(name string) bool {
	_, ok := db.names[strings.ToLower(name)]
	return ok
}

// HasID checks if a user with the given id exists in the
// database.
func (db *UserDB) HasID(id uint32) bool {
	db.RLock()
	defer db.RUnlock()
	return db.hasID(id)
}

func (db *UserDB) hasID(id uint32) bool {
	_, ok := db.ids[id]
	return ok
}

// ByName fetches a user from the database that matches the
// username given. returns nil if none was found.
func (db *UserDB) ByName(name string) (User, error) {
	db.RLock()
	defer db.RUnlock()
	if u := db.names[strings.ToLower(name)]; u != nil {
		return u.Copy(), nil
	}
	return nil, &ErrNoUser{}
}

// ByID fetches a user from the database that matches the
// id given. returns nil if none was found.
func (db *UserDB) ByID(id uint32) (User, error) {
	db.RLock()
	defer db.RUnlock()
	if u := db.ids[id]; u != nil {
		return u.Copy(), nil
	}
	return nil, &ErrNoUser{}
}

// Remove deletes the user from the database.
func (db *UserDB) Remove(u User) {
	db.Lock()
	db.removeID(u.ID())
	db.Unlock()
}

// RemoveName deletes the user that matches the given name
// from the database.
func (db *UserDB) RemoveName(name string) {
	db.Lock()
	db.removeName(name)
	db.Unlock()
}

func (db *UserDB) removeName(name string) {
	if db.hasName(name) {
		name = strings.ToLower(name)
		id := db.names[name].ID()
		delete(db.names, name)
		delete(db.ids, id)
	}
}

// RemoveID deletes the user that matches the given id from
// the database.
func (db *UserDB) RemoveID(id uint32) {
	db.Lock()
	db.removeID(id)
	db.Unlock()
}

func (db *UserDB) removeID(id uint32) {
	if db.hasID(id) {
		name := strings.ToLower(db.ids[id].Name())
		delete(db.names, name)
		delete(db.ids, id)
	}
}

// Verify returns a boolean value indicating whether user
// and password match entries in the user database.
func Verify(pass, hash string) bool {
	if pass == "" || hash == "" {
		return false
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil {
		return true
	}
	return false
}

// Encrypt generates a cryptographically secure hash from the
// given string. It is strong enough for password storage.
func Encrypt(s string) string {
	if k, err := bcrypt.GenerateFromPassword([]byte(s), 12); err == nil {
		return fmt.Sprintf("%s\n", k)
	}
	return ""
}

// ErrNoUser indicates that a user was not found in the cache.
type ErrNoUser struct{}

func (e *ErrNoUser) Error() string {
	return "user not found"
}
