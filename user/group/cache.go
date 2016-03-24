package group

import "sail/user/rights"

type cache struct {
	groups map[uint32]*Group
}

var cacheInstance *cache

// Cache contains all currently known user groups in the system.
func Cache() *cache {
	if cacheInstance == nil {
		gs := make(map[uint32]*Group)
		for _, g := range fromStorageByID() {
			gs[g.ID] = g
		}
		cacheInstance = &cache{groups: gs}
	}
	return cacheInstance
}

// At returns the group with the given id if it exists, nil otherwise.
func (c *cache) At(id uint32) *Group {
	return c.groups[id]
}

// Permission returns a user's permission object relating to the
// given permission domain.
func (c *cache) Permission(uid uint32, dom rights.Domain) *rights.Permission {
	return rights.New(dom, c.Mode(uid, dom))
}

// Mode returns the access mode granted to the user in relation to
// the given permission domain.
func (c *cache) Mode(uid uint32, dom rights.Domain) (m rights.Mode) {
	for _, g := range c.groups {
		if _, ok := g.users[uid]; ok {
			m = m | g.Mode(dom)
			if m == 7 {
				break
			}
		}
	}
	return
}
