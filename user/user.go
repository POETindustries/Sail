package user

import "errors"

// User represents a Person registered within the system.
type User struct {
	id        uint32
	name      string
	pass      string
	FirstName string
	LastName  string
	Email     string
	Phone     string

	CDate   string
	ExpDate string
}

// New returns a fresh, minimally initialized User object.
func New(name string) *User {
	return &User{name: name}
}

// TODO 2017-02-08: rename to LoadOrNew
func LoadNew(name string) *User {
	u := New(name)
	if u.Load() {
		return u
	}
	return New(name)
}

func (u *User) ID() uint32 {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Load() bool {
	return singleFromStorage(u)
}

func (u *User) Hash() (string, error) {
	if u.pass == "" {
		return "", errors.New("No hash found")
	}
	return u.pass, nil
}

func (u *User) Save() bool {
	return false
}

func (u *User) SaveHash(hash string) bool {
	return false
}

func (u *User) Store() error {
	return nil
}
