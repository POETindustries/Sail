package group

import "sail/user/permission"

type Group struct {
	ID   uint32
	Name string
	perm [2]permission.Mode
}

func (g *Group) Permission(domain permission.Domain) *permission.Permission {
	return permission.New(domain, g.perm[domain])
}
