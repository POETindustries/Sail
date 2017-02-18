package group

import "sail/user/rights"

// Cache manages an internal pool of groups.
type Cache struct {
	groups map[uint32]*Group
}

var cache *Cache

// All contains all currently known user groups in the system.
func All() *Cache {
	if cache == nil {
		gs := make(map[uint32]*Group)
		for _, g := range fromStorageByID() {
			gs[g.ID] = g
		}
		cache = &Cache{groups: gs}
	}
	return cache
}

// At returns the group with the given id if it exists, nil otherwise.
func (c *Cache) At(id uint32) *Group {
	return c.groups[id]
}

// Permission returns a user's permission object relating to the
// given permission domain.
func (c *Cache) Permission(uid uint32, dom rights.Domain) *rights.Permission {
	return rights.New(dom, c.Mode(uid, dom))
}

// Mode returns the access mode granted to the user in relation to
// the given permission domain.
func (c *Cache) Mode(uid uint32, dom rights.Domain) (m rights.Mode) {
	for _, g := range c.groups {
		if _, ok := g.users[uid]; ok {
			m = m | g.Mode(dom)
			if m == 15 {
				break
			}
		}
	}
	return
}
