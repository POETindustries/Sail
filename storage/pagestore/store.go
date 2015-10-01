package pagestore

import (
	"bytes"
	"sail/errors"
	"sail/page"
	"sail/storage/conn"
	"strconv"
)

// Query collects all information needed for querying the database.
type Query struct {
	table     string
	proj      string
	sel       bytes.Buffer
	selAttrs  []interface{}
	attrCount int
}

// Visible prepares the query to select all pages that are visible
// and accessible from the general internet.
func (q *Query) Visible() *Query {
	q.addAttr(pageStatus, -1)
	return q
}

// ByURL prepares the query to select only those pages that match the
// given url string.
func (q *Query) ByURL(url string) *Query {
	q.addAttr(pageURL, url)
	return q
}

// ByID prepares the query to select the page that matches the given id.
func (q *Query) ByID(id uint32) *Query {
	q.addAttr(pageID, id)
	return q
}

// Execute sends the query to the database and writes the resulting
// data to the Page parameter.
//
// Any error that is raised during execution is returned. An error is
// also returned if the query's selection does not contain at least
// one attribute.
func (q *Query) Execute(p *page.Page) error {
	if q.sel.Len() < 1 {
		return errors.NoArguments()
	}
	data := conn.Instance().DB.QueryRow(q.build(), q.selAttrs...)

	return data.Scan(&p.ID, &p.Title, &p.Content, &p.Domain.ID, &p.URL,
		&p.Status, &p.Owner, &p.CDate, &p.EDate)
}

func (q *Query) addAttr(key string, val interface{}) {
	if q.attrCount > 1 {
		q.sel.WriteString(" and ")
	}
	q.sel.WriteString(key + "=$" + strconv.Itoa(q.attrCount))
	q.selAttrs = append(q.selAttrs, val)
	q.attrCount++
}

func (q *Query) build() string {
	return "select " + q.proj + " from " + q.table + " where " + q.sel.String()
}

// Get starts building the query that gets sent to the database.
//
// TODO: describe how queries should be built using method chaining.
func Get() *Query {
	return &Query{attrCount: 1, table: "sl_page", proj: pageAttrs}
}
