package user

type Manager struct {
	users []*User
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) All() []*User {
	if m.users == nil {
		m.users = fromStorageByName()
	}
	return m.users
}
