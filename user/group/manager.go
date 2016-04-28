package group

import "sail/user"

type Manager struct {
	userMgr *user.Manager
	groups  []*Group
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) LoadUsers() {
	if m.userMgr == nil {
		m.userMgr = user.NewManager()
	}
	m.userMgr.All()
}

func (m *Manager) AllUsers() []*user.User {
	if m.userMgr == nil {
		m.LoadUsers()
	}
	return m.userMgr.All()
}

func (m *Manager) AllGroups() []*Group {
	if m.groups == nil {
		m.groups = m.fromCache()
	}
	return m.groups
}

func (m *Manager) fromCache() []*Group {
	var gs []*Group
	for _, g := range All().groups {
		gs = append(gs, g)
	}
	return gs
}
