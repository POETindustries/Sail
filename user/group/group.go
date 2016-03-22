package group

import "sail/user/rights"

type Group struct {
	ID   uint32
	Name string

	users map[uint32]bool
	perm  [2]rights.Mode
}

func New() *Group {
	return &Group{users: make(map[uint32]bool)}
}

func (g *Group) Permission(dom rights.Domain) *rights.Permission {
	return rights.New(dom, g.Mode(dom))
}

func (g *Group) Mode(dom rights.Domain) rights.Mode {
	if dom < rights.DomainCount {
		return g.perm[dom]
	}
	return 0
}
