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

package input

import "regexp"

var (
	alphanum = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	alpha    = regexp.MustCompile(`^[a-zA-Z]+$`)
	num      = regexp.MustCompile(`^[0-9]+$`)

	email           = regexp.MustCompile(`^.+@.+$`)
	username        = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_-]{1,22}[a-zA-Z0-9]$`)
	badPassTopology = []*regexp.Regexp{
		// contains only letters
		alpha,
		// contains only numbers
		num,
		// empty string
		regexp.MustCompile(`^$`),
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

// IsValidUsername returns true if s is a valid username according to
// the following guidelines:
//
// - contains only printable alphanumeric ASCII characters plus
//   hyphen and underscore,
// - cannot start or end with hyphen or underscore,
// - is at least 3 and at most 24 characters long
//
// Some conditions may seem arbitrary or too restricting. They are
// chosen with the role of usernames as uniquely identifying users
// to other users and machines in mind.
//
// Limiting the character range to ASCII makes processing by software
// easier, especially when checkign equality and the like. Since
// usernames can be part of URLs and other resource identifiers,
// having them withing ASCII range ensures some amount of backwards
// compatibility with software that doesn't handle unicode well.
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
