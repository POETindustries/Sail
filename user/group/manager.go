package group

import "sail/session"

type Manager struct {
	userMgr *session.UserDB
	groups  []*Group
}

func NewManager() *Manager {
	return &Manager{userMgr: session.Users()}
}

func (m *Manager) LoadUsers() {
	if m.userMgr == nil {
		m.userMgr = session.Users()
	}
}

func (m *Manager) AllUsers() []session.User {
	// TODO 2017-02-03: returns only currently active
	// users, should fetch from persistent storage.
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
