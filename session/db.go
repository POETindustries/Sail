package session

type DB struct {
	sessions map[string]*Session
}

var instance *DB

func new() *DB {
	db := DB{}
	db.sessions = make(map[string]*Session)

	return &db
}

func DBInstance() *DB {
	if instance == nil {
		instance = new()
	}
	return instance
}

func (db *DB) Add(session *Session) {
	db.sessions[session.ID] = session
}

func (db *DB) Remove(id string) {
	delete(db.sessions, id)
}

func (db *DB) ID(id string) string {
	return db.sessions[id].ID
}

func (db *DB) User(id string) string {
	return db.sessions[id].User
}

func (db *DB) Lang(id string) string {
	return db.sessions[id].Lang
}

func (db *DB) Has(id string) bool {
	return DBInstance().sessions[id] != nil
}
