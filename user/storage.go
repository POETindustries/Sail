package user

import (
	"database/sql"
	"sail/user/schema"
)

// Query collects all information needed for querying the database.
type Query struct {
	query *psqldb.Query
}

// ByID prepares the query to select the pages that matches the id(s).
func (q *Query) ByID(ids ...uint32) *Query {
	for _, id := range ids {
		q.query.AddAttr(schema.UserID, id, psqldb.OpOr)
	}
	return q
}

func (q *Query) ByName(name string) *Query {
	q.query.AddAttr(schema.UserName, name, psqldb.OpAnd)
	return q
}

// Users sends the query to the database and returns all matching
// user objects.
func (q *Query) Users() ([]*data.User, error) {
	q.query.Table = "sl_user"
	q.query.Proj = schema.UserAttrs
	return q.scanUsers(q.query.Execute())
}

func (q *Query) scanUsers(rows *sql.Rows, err error) ([]*data.User, error) {
	if err != nil {
		return nil, err
	}
	users := []*data.User{}
	defer rows.Close()
	for rows.Next() {
		u := data.New()
		if err = rows.Scan(&u.ID, &u.Name, &u.Pass, &u.FirstName, &u.LastName,
			&u.Email, &u.Phone, &u.CDate, &u.ExpDate); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Get starts building the query that gets sent to the database.
//
// TODO: describe how queries should be built using method chaining.
func Get() *Query {
	return &Query{query: &psqldb.Query{}}
}
