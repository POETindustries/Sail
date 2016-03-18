package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

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

func New() *User {
	return &User{}
}

func ByName(name string) *User {
	us := fromStorageByName(name)
	if len(us) < 1 {
		return nil
	}
	return us[0]
}

// Verify returns true if user and password match entries in the
// user database.
func Verify(user, pass string) bool {
	u := ByName(user)
	if u == nil {
		return false
	}
	p := []byte(pass)
	h := []byte(u.pass)
	return bcrypt.CompareHashAndPassword(h, p) == nil
}

func encrypt(s string) string {
	if k, err := bcrypt.GenerateFromPassword([]byte(s), 12); err == nil {
		return fmt.Sprintf("%s\n", k)
	}
	return ""
}
