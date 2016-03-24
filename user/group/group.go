package group

import "sail/user/rights"

// Group represents a user group within the system. It manages a list
// of user ids from all its members and regulates file access and
// query action permissions for all users that belong to the group.
type Group struct {
	ID   uint32
	Name string

	users map[uint32]bool
	perm  [2]rights.Mode
}

// New returns an empty Group object.
func New() *Group {
	return &Group{users: make(map[uint32]bool)}
}

// Permission returns the group's permission object relating to the
// given permission domain.
func (g *Group) Permission(dom rights.Domain) *rights.Permission {
	return rights.New(dom, g.Mode(dom))
}

// Mode returns the access mode granted to the group in relation to
// the given permission domain.
func (g *Group) Mode(dom rights.Domain) rights.Mode {
	if dom < rights.DomainCount {
		return g.perm[dom]
	}
	return 0
}
