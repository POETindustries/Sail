package pagestore

import (
	"database/sql"
	"html/template"
	"sail/page"
	"sail/storage/psqldb"
)

// Query collects all information needed for querying the database.
type Query struct {
	query *psqldb.Query
}

// Visible prepares the query to select all pages that are visible
// and accessible from the general internet.
func (q *Query) Visible() *Query {
	q.query.AddAttr(pageStatus, -1, "")
	return q
}

// ByURL prepares the query to select only those pages that match the
// given url(s).
func (q *Query) ByURL(urls ...string) *Query {
	for _, url := range urls {
		q.query.AddAttr(pageURL, url, "")
	}
	return q
}

// ByID prepares the query to select the pages that matches the id(s).
func (q *Query) ByID(ids ...uint32) *Query {
	for _, id := range ids {
		q.query.AddAttr(pageID, id, psqldb.OpOr)
	}
	return q
}

// Pages sends the query to the database and returns all matching
// page objects.
func (q *Query) Pages() ([]*page.Page, error) {
	q.query.Table = "sl_page"
	q.query.Proj = pageAttrs
	return q.scanPages(q.query.Execute())
}

func (q *Query) scanPages(data *sql.Rows, err error) ([]*page.Page, error) {
	if err != nil {
		return nil, err
	}
	pages := []*page.Page{}
	defer data.Close()
	for data.Next() {
		var c string
		p := page.New()
		if err = data.Scan(&p.ID, &p.Title, &c, &p.Domain.ID, &p.URL,
			&p.Status, &p.Owner, &p.CDate, &p.EDate); err != nil {
			return nil, err
		}
		p.Content = template.HTML(c)
		pages = append(pages, p)
	}
	return pages, nil
}

// Get starts building the query that gets sent to the database.
//
// TODO: describe how queries should be built using method chaining.
func Get() *Query {
	return &Query{query: &psqldb.Query{}}
}
