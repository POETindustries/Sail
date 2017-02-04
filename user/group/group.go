package group

import "sail/user/rights"

// Group represents a user group within the system. It manages a list
// of user ids from all its members and regulates file access and
// query action permissions for all users that belong to the group.
type Group struct {
	ID   uint32
	Name string

	users map[uint32]bool
	perm  [4]rights.Mode
}

// New returns an empty Group object.
func New() *Group {
	return &Group{users: make(map[uint32]bool)}
}

// MemberCount returns how many members the group currently has.
func (g *Group) MemberCount() int {
	return len(g.users)
}

// Mode returns the access mode granted to the group in relation to
// the given permission domain.
func (g *Group) Mode(dom rights.Domain) rights.Mode {
	if dom < rights.DomainCount {
		return g.perm[dom]
	}
	return 0
}

// Permission returns the group's permission object relating to the
// given permission domain.
func (g *Group) Permission(dom rights.Domain) *rights.Permission {
	return rights.New(dom, g.Mode(dom))
}

// PermissionStatus gives an Impression of the group's overall
// privileges, indicated by a number ranging from 0 (least
// privileges) to 100 (permission to do anything anywhere).
func (g *Group) PermissionStatus() int {
	stat := 0
	for _, p := range g.perm {
		stat += int(p)
	}
	return (stat * 100) / (15 * len(g.perm))
}
