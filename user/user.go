package user

import (
	"fmt"
	"sail/storage/userstore"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32
	Name      string
	Pass      string
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

// Verify returns true if user and password match entries in the
// user database.
func Verify(user, pass string) bool {
	users, err := userstore.Get().ByName(user).Users()
	if err != nil || len(users) == 0 {
		return false
	}
	p := []byte(pass)
	h := []byte(users[0].Pass)
	return bcrypt.CompareHashAndPassword(h, p) == nil
}

func encrypt(s string) string {
	if k, err := bcrypt.GenerateFromPassword([]byte(s), 8); err == nil {
		return fmt.Sprintf("%s\n", k)
	}
	return ""
}
