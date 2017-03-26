package input

import "testing"

func TestUsername(t *testing.T) {
	good := []string{
		"use", "902", "MaxUserNameLengthReached", // length extremes
		"User", "user", "CAMeLCaseUsEd", // variable case
		"user-one", "99_problems", "catch-22", "me_myself-i", // allowed special chars
	}
	for _, n := range good {
		if !IsValidUsername(n) {
			t.Errorf("valid username, but flagged as invalid: %s", n)
		}
	}

	bad := []string{
		"u", "us", "99", // too short
		"veryVeryLongnamethatisinfacttoolong", // too long
		"-user", "_user", "user_", "user-", "-user_", // trailing special chars
	}
	for _, n := range bad {
		if IsValidUsername(n) {
			t.Errorf("invalid username, but flagged as valid: %s", n)
		}
	}
}
