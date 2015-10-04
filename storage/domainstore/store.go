package domainstore

import (
	"database/sql"
	"sail/domain"
	"sail/storage/psqldb"
)

// Query collects all information needed for querying the database.
type Query struct {
	query *psqldb.Query
}

// ByID prepares the query to select the domain that matches the id(s).
func (q *Query) ByID(ids ...uint32) *Query {
	for _, id := range ids {
		q.query.AddAttr(domainID, id, psqldb.OpOr)
	}
	return q
}

// Domains executes the query and returns all matching domain objects.
func (q *Query) Domains() ([]*domain.Domain, error) {
	q.query.Table = "sl_domain"
	q.query.Proj = domainAttrs
	return q.scanDomains(q.query.Execute())
}

func (q *Query) scanDomains(data *sql.Rows, err error) ([]*domain.Domain, error) {
	if err != nil {
		return nil, err
	}
	domains := []*domain.Domain{}
	defer data.Close()
	for data.Next() {
		d := domain.New()
		if err = data.Scan(&d.ID, &d.Name, &d.Meta.Title, &d.Meta.Keywords,
			&d.Meta.Description, &d.Meta.Language, &d.Meta.PageTopic,
			&d.Meta.RevisitAfter, &d.Meta.Robots, &d.Template.ID); err != nil {
			return nil, err
		}
		domains = append(domains, d)
	}
	return domains, nil
}

// Get starts building the query that gets sent to the database.
//
// TODO: describe how queries should be built using method chaining.
func Get() *Query {
	return &Query{query: &psqldb.Query{}}
}
