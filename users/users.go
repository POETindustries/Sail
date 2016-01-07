package users

import (
	"fmt"
	"sail/storage/userstore"

	"golang.org/x/crypto/bcrypt"
)

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
