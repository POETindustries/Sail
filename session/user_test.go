package session

import (
	"strings"
	"testing"
)

//////////////////////////////////////////////////////////////////////
// Mockup Data Structures & Objects
/////////////////////////////////////////////////////////////////////

type testuser struct {
	id   uint32
	name string
	age  uint8
}

func (u *testuser) ID() uint32 {
	return u.id
}

func (u *testuser) Name() string {
	return u.name
}

func (u *testuser) Copy() User {
	return &testuser{u.id, u.name, u.age}
}

var testuser1 = &testuser{id: 1, name: "TestUser1", age: 12}
var testuser2 = &testuser{id: 2, name: "TestUser2", age: 24}
var testuser3 = &testuser{id: 3, name: "TestUser3", age: 36}

//////////////////////////////////////////////////////////////////////
// Begin Tests
/////////////////////////////////////////////////////////////////////

func TestAddUser(t *testing.T) {
	db := NewUserDB()

	// should succeed
	db.Add(testuser1)
	if len(db.names) != 1 || len(db.ids) != 1 {
		t.Error("add user failed")
	} else if db.ids[testuser1.ID()] != db.names[strings.ToLower(testuser1.Name())] {
		t.Error("unintended data duplication")
	} else if db.ids[testuser1.ID()] == testuser1 {
		t.Error("unexpected object identity")
	}

	// should succeed
	db.Add(testuser1)
	if len(db.names) != 1 || len(db.ids) != 1 {
		t.Error("add user failed")
	} else if db.ids[testuser1.ID()] != db.names[strings.ToLower(testuser1.Name())] {
		t.Error("unintended data duplication")
	} else if db.ids[testuser1.ID()] == testuser1 {
		t.Error("unexpected object identity")
	}

	// should succeed
	db.Add(testuser2)
	if len(db.names) != 2 || len(db.ids) != 2 {
		t.Error("add user failed")
	} else if db.ids[testuser2.ID()] != db.names[strings.ToLower(testuser2.Name())] {
		t.Error("unintended data duplication")
	} else if db.ids[testuser2.ID()] == testuser2 {
		t.Error("unexpected object identity")
	}

	// should succeed
	db.Add(testuser3)
	if len(db.names) != 3 || len(db.ids) != 3 {
		t.Error("add user failed")
	} else if db.ids[testuser3.ID()] != db.names[strings.ToLower(testuser3.Name())] {
		t.Error("unintended data duplication")
	} else if db.ids[testuser3.ID()] == testuser3 {
		t.Error("unexpected object identity")
	}
}

func TestHasUser(t *testing.T) {
	db := NewUserDB()

	// should succeed
	db.Add(testuser1)
	if !db.Has(testuser1) {
		t.Error("should have failed")
	} else if !db.HasID(testuser1.ID()) {
		t.Error("should have failed")
	} else if !db.HasName(testuser1.Name()) {
		t.Error("should have failed")
	} else if !db.HasName(strings.ToLower(testuser1.Name())) {
		t.Error("should have failed")
	} else if !db.HasName(strings.ToUpper(testuser1.Name())) {
		t.Error("should have failed")
	}

	// should fail, user not in cache
	if db.Has(testuser2) {
		t.Error("should have failed")
	} else if db.HasID(testuser2.ID()) {
		t.Error("should have failed")
	} else if db.HasName(testuser2.Name()) {
		t.Error("should have failed")
	}

	// should succeed
	db.Add(testuser2)
	if !db.Has(testuser2) {
		t.Error("should have failed")
	} else if !db.HasID(testuser2.ID()) {
		t.Error("should have failed")
	} else if !db.HasName(testuser2.Name()) {
		t.Error("should have failed")
	}
}

func TestGetUser(t *testing.T) {
	db := NewUserDB()

	// should succeed and return copy of user
	db.Add(testuser1)
	if u, err := db.ByID(testuser1.ID()); err != nil {
		t.Error(err)
	} else if u.ID() != testuser1.ID() {
		t.Error("invalid data")
	} else if u == testuser1 {
		t.Error("unexpected object identity")
	} else if u2, err := db.ByName(testuser1.Name()); err != nil {
		t.Error(err)
	} else if u2.ID() != testuser1.ID() {
		t.Error("invalid data")
	} else if u == u2 {
		t.Error("unexpected object identity")
	}

	// should fail, user is not in cache
	if _, err := db.ByID(testuser2.ID()); err == nil {
		t.Error("should have failed")
	} else if _, ok := err.(*ErrNoUser); !ok {
		t.Error("wrong error:", err.Error())
	}

	// should succeed and return copy of user
	db.Add(testuser2)
	if u, err := db.ByID(testuser2.ID()); err != nil {
		t.Error(err)
	} else if u.ID() != testuser2.ID() {
		t.Error("invalid data")
	} else if u == testuser2 {
		t.Error("unexpected object identity")
	} else if u2, err := db.ByName(testuser2.Name()); err != nil {
		t.Error(err)
	} else if u2.ID() != testuser2.ID() {
		t.Error("invalid data")
	} else if u == u2 {
		t.Error("unexpected object identity")
	}
}

func TestAllUsers(t *testing.T) {
	db := NewUserDB()

	// should have no users in it
	if len(db.All()) != 0 {
		t.Error("unexpected data")
	}

	// should return all usernames
	db.Add(testuser1)
	db.Add(testuser2)
	db.Add(testuser3)
	if a := db.All(); len(a) != 3 {
		t.Error("missing user")
	} else if a[0] != testuser1.Name() && a[0] != testuser2.Name() && a[0] != testuser3.Name() {
		t.Error("invalid data modification")
	}
}

func TestRemoveUser(t *testing.T) {
	db := NewUserDB()

	// should succeed but not remove anything, cache is empty
	db.Remove(testuser1)
	if len(db.ids) != len(db.names) || len(db.ids) != 0 {
		t.Error("unexpected cache contents")
	}

	// should succeed, but not remove anything, user is not in cache
	db.Add(testuser1)
	db.Remove(testuser2)
	if len(db.ids) != len(db.names) || len(db.ids) != 1 {
		t.Error("unexpected cache contents")
	} else if !db.Has(testuser1) {
		t.Error("unintentional data modification")
	}

	// should succeed and remove user from cache
	db.Add(testuser2)
	db.Remove(testuser1)
	if len(db.ids) != len(db.names) || len(db.ids) != 1 {
		t.Error("unexpected cache contents")
	} else if !db.Has(testuser2) {
		t.Error("unintentional data modification")
	} else if db.Has(testuser1) {
		t.Error("unexpected cache contents")
	}
}
