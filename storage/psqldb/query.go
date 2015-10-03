package psqldb

import (
	"bytes"
	"database/sql"
	"fmt"
	"sail/errors"
	"strconv"
)

// OpAnd is the statement-ready and operator for query building.
const OpAnd = " and "

// OpOr is the statement-ready or operator for query building.
const OpOr = " or "

// Query collects all information needed for querying the database.
type Query struct {
	Table     string
	Proj      string
	sel       bytes.Buffer
	selAttrs  []interface{}
	attrCount int

	// -1: desc // 0: not set // +1: asc
	order     int8
	orderStmt string
}

// AddAttr adds the attribute, specified by the key-val pair, to the
// selection string. An operator can be passed for specifiying how
// the selection attributes should be connected, if there are more
// than one of them.
func (q *Query) AddAttr(key string, val interface{}, op string) {
	q.attrCount++
	if q.attrCount > 1 {
		q.sel.WriteString(op)
	}
	q.sel.WriteString(key + "=$" + strconv.Itoa(q.attrCount))
	q.selAttrs = append(q.selAttrs, val)
}

// Execute builds the sql query from the current values of its fields
// and queries the database.
func (q *Query) Execute() (*sql.Rows, error) {
	if q.sel.Len() < 1 {
		return nil, errors.NoArguments()
	}
	return Instance().DB.Query(q.build(), q.selAttrs...)
}

// Ascending sets the appropriate order flag.
func (q *Query) Ascending() { q.order = 1 }

// Descending sets the appropriate order flag.
func (q *Query) Descending() { q.order = -1 }

// Order returns the projection order in a statement-ready format.
// An empty string is returned when no order is set.
func (q *Query) Order() string {
	if q.order == 1 {
		return " asc "
	} else if q.order == -1 {
		return " desc "
	}
	return ""
}

// SetOrderStmt passes the complete sql order statement to Query.
func (q *Query) SetOrderStmt(stmt string) {
	q.orderStmt = stmt
}

func (q *Query) build() string {
	query := "select %s from %s where %s %s"
	return fmt.Sprintf(query, q.Proj, q.Table, q.sel.String(), q.orderStmt)
}
