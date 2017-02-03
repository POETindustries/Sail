package session

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User interface {
	ID() uint32
	Name() string
	Load() bool
	Hash() (string, error)
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
