package user

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
