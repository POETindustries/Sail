package group

import "sail/user/permission"

type Group struct {
	ID   uint32
	Name string

	users map[uint32]bool
	perm  [2]permission.Mode
}

func New() *Group {
	return &Group{users: make(map[uint32]bool)}
}

func (g *Group) Permission(domain permission.Domain) *permission.Permission {
	return permission.New(domain, g.Mode(domain))
}

func (g *Group) Mode(domain permission.Domain) permission.Mode {
	if domain < permission.DomainCount {
		return g.perm[domain]
	}
	return 0
}
