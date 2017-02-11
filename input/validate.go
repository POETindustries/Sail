package input

import "regexp"

var (
	alphanum = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	alpha    = regexp.MustCompile(`^[a-zA-Z]+$`)
	num      = regexp.MustCompile(`^[0-9]+$`)

	email           = regexp.MustCompile(`^.+@.+$`)
	username        = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,24}$`)
	badPassTopology = []*regexp.Regexp{
		// contains only letters
		alpha,
		// contains only numbers
		num,
		// less than 8 characters
		regexp.MustCompile(`^.{1,7}$`),
		// Topology: Password1, passwords123, ...
		regexp.MustCompile(`^[A-Za-z][a-z]{6,10}[0-9]{1,4}$`)}
)

func IsAlphanum(s string) bool {
	return alphanum.MatchString(s)
}

func IsAlpha(s string) bool {
	return alpha.MatchString(s)
}

func IsNum(s string) bool {
	return num.MatchString(s)
}

func IsEmail(s string) bool {
	return email.MatchString(s)
}

func IsValidUsername(s string) bool {
	return username.MatchString(s)
}

func IsValidPassword(s string) bool {
	for _, r := range badPassTopology {
		if r.MatchString(s) {
			return false
		}
	}
	return true
}
