package session

type Database struct {
	sessions map[string]*Session
}

var instance *Database

func new() *Database {
	db := Database{}
	db.sessions = make(map[string]*Session)

	return &db
}

func DB() *Database {
	if instance == nil {
		instance = new()
	}
	return instance
}

func (db *Database) Add(session *Session) {
	db.sessions[session.ID] = session
}

func (db *Database) Remove(id string) {
	delete(db.sessions, id)
}

func (db *Database) ID(id string) string {
	return db.sessions[id].ID
}

func (db *Database) User(id string) string {
	return db.sessions[id].User
}

func (db *Database) Lang(id string) string {
	return db.sessions[id].Lang
}

func (db *Database) Has(id string) bool {
	return DB().sessions[id] != nil
}
