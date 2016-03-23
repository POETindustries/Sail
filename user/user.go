package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User represents a Person registered within the system.
type User struct {
	ID        uint32
	Name      string
	pass      string
	FirstName string
	LastName  string
	Email     string
	Phone     string

	CDate   string
	ExpDate string
}

// New returns a fresh, minimally initialized User object.
func New() *User {
	return &User{}
}

// ByName returns a User object from the persistent storage,
// identified by its unique user name.
func ByName(name string) *User {
	us := fromStorageByName(name)
	if len(us) < 1 {
		return nil
	}
	return us[0]
}

// Verify returns a User and a boolean value indicating whether user
// and password match entries in the user database. If so, the User
// return value is safe for use.
//
// Since there has to be a database request anyway, the resulting
// User object is returned in order to give the opportunity to use it
// right away instead of fetching anew a few lines later.
func Verify(user, pass string) (u *User, ok bool) {
	if u = ByName(user); u != nil {
		ok = bcrypt.CompareHashAndPassword([]byte(u.pass), []byte(pass)) == nil
	}
	return
}

func encrypt(s string) string {
	if k, err := bcrypt.GenerateFromPassword([]byte(s), 12); err == nil {
		return fmt.Sprintf("%s\n", k)
	}
	return ""
}
