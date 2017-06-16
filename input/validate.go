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
		// Topology: Password1, passwords123, NewPassword1!, ...
		regexp.MustCompile(`^[A-Za-z][a-z]{6,10}[0-9]{0,4}[!\?]{0,1}$`)}
)

// IsAlphanum returns true when a string consists of only letters and numbers.
func IsAlphanum(s string) bool {
	return alphanum.MatchString(s)
}

// IsAlpha returns true if a string consists only of letters.
func IsAlpha(s string) bool {
	return alpha.MatchString(s)
}

// IsNum is true if a string contains only numbers.
func IsNum(s string) bool {
	return num.MatchString(s)
}

// IsEmail is true if the string is a valid e-mail address. Note that
// this check is very permissive in order to prevent false negatives.
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

// IsValidPassword returns true if a string is suitable for being used
// as a password that satisfies a set of security conditions.
//
// IsValidPassword checks not only against commonly used password
// security guidelines, like minimum length and the use of different
// character classes. It also checks for common topologies to prevent
// a password's structure from reducing the effort of guessing it. See
// https://www.korelogic.com/Resources/Presentations/bsidesavl_pathwell_2014-06.pdf
// for more info on the topic.
func IsValidPassword(s string) bool {
	for _, r := range badPassTopology {
		if r.MatchString(s) {
			return false
		}
	}
	return true
}
