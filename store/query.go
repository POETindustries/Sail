package store

import (
	"database/sql"
	"fmt"
	"sail/conf"
	"sail/errors"
	"strings"
)

const (
	and  = "and"
	or   = "or"
	asc  = "asc"
	desc = "desc"
)

const (
	modeGet    = 0x001
	modeAdd    = 0x002
	modeUpdate = 0x004
	modeDelete = 0x008
)

// Query acts as middle man between persistent storage and
// business logic. All requests to persistent storage should
// happen through query objects.
type Query struct {
	table         string
	attrs         []string
	attrVals      []interface{}
	selection     []string
	selectionVals []interface{}
	order         string
	orderAttr     string

	// we need a copy of the driver because we
	// cannot guarantee that its state won't be
	// changed by other goroutine running parallel
	// to this one.
	driver Driver
	mode   uint8
}

// Add sets the query's operation mode to insertion, signaling
// that the data being sent is to be inserted as a new dataset.
func Add() *Query {
	return &Query{mode: modeAdd, driver: DB().driver.Copy()}
}

// Delete sets the query's operation mode to deletion. The
// dataset is to be deleted from permanent storage.
func Delete() *Query {
	return &Query{mode: modeDelete, driver: DB().driver.Copy()}
}

// Get indicates that data is to be retrieved from storage,
// not changing any values.
func Get() *Query {
	return &Query{mode: modeGet, driver: DB().driver.Copy()}
}

// Update sends data to update datasets that are expected to
// already exist in the database.
func Update() *Query {
	return &Query{mode: modeUpdate, driver: DB().driver.Copy()}
}

// All indicates that all datasets from a given table should
// be targeted. Ignored by operation mode Add().
// 		Danger Zone: This is potentially harmful, especially
//		in combination with Delete(). Use of All() should always
//		go with extra attention.
func (q *Query) All() *Query {
	q.selection = append(q.selection, "1=1")
	return q
}

// And adds an additional condition to the query's selection.
// Its position in the instruction chain is meaningful.
func (q *Query) And() *Query {
	if len(q.selection) > 0 {
		q.selection = append(q.selection, and)
	}
	return q
}

// Asc instructs query to order the results in ascending order,
// i.e. from lowest to highest.
func (q *Query) Asc() *Query {
	q.order = asc
	return q
}

// Attrs is used to set which attributes, or columns, to use
// when executing the query. Their usage depends on the query's
// operation mode. Get() retrieves the attributes, while all
// other modes ignore them completely.
func (q *Query) Attrs(attrs ...string) *Query {
	if q.mode == modeGet {
		q.attrs = attrs
	}
	return q
}

// Desc instructs query to order the results in descending
// order, i.e. from highest to lowest.
func (q *Query) Desc() *Query {
	q.order = desc
	return q
}

// Equals passes the key-value pair relevant for matching the
// datasets to the query. It instructs to work only on those
// datasets where the value of attribute 'key' matches the
// value passed with 'val'.
func (q *Query) Equals(key string, val interface{}) *Query {
	q.addSelection(key, val, "="+q.driver.Param())
	return q
}

// Exec executes the query, making the actual request to the
// database. It should be the last operation that query does
// before retrieving the results.
func (q *Query) Exec() (rows *sql.Rows, ok bool) {
	var err error
	switch q.mode {
	case modeGet, modeDelete:
		rows, err = DB().Query(q.build(), q.selectionVals...)
	case modeAdd, modeUpdate:
		vals := q.driver.PrepareData(q)
		_, err = DB().Exec(q.build(), vals...)
	}
	ok = (err == nil)
	if !ok {
		errors.Log(err, conf.Instance().DevMode)
	}
	return
}

// In is used to pass the table that query should be prepared for.
func (q *Query) In(table string) *Query {
	q.table = table
	return q
}

// NotEquals acts as opposite of Equals(). It passes the key-value
// pair relevant for matching the datasets to the query and
// instructs to work only on those datasets where the value of
// attribute 'key' does not match the value passed with 'val'.
func (q *Query) NotEquals(key string, val interface{}) *Query {
	q.addSelection(key, val, "<>"+q.driver.Param())
	return q
}

// Or adds another condition to the query.
func (q *Query) Or() *Query {
	if len(q.selection) > 0 {
		q.selection = append(q.selection, or)
	}
	return q
}

// Order is used to define the attribute that should be the
// one that determines sort order. Only relevant when used
// with Asc() or Desc().
func (q *Query) Order(attr string) *Query {
	q.orderAttr = attr
	return q
}

// String implements the Printable interface and prints a
// human-readable representation of the query if it would
// be executed in the state it is at the current time.
func (q *Query) String() string {
	copy := Query{
		mode:      q.mode,
		table:     q.table,
		order:     q.order,
		orderAttr: q.orderAttr,
		driver:    q.driver.Copy()}
	copy.attrs = append(copy.attrs, q.attrs...)
	copy.attrVals = append(copy.attrVals, q.attrVals...)
	copy.selection = append(copy.selection, q.selection...)
	copy.selectionVals = append(copy.selectionVals, q.selectionVals...)
	vals := copy.driver.PrepareData(&copy)
	return fmt.Sprintf("%s|%v\n", copy.build(), vals)
}

// Values is used to pass the 'payload' to the query. It takes
// a set of key-value pairs that are handled based on the
// operation mode. When updating or inserting a dataset, vals
// contain the new values to write. Get() and Delete() ignore
// these completely.
func (q *Query) Values(vals map[string]interface{}) *Query {
	if q.mode == modeUpdate {
		for k, v := range vals {
			q.attrs = append(q.attrs, k+"="+q.driver.Param())
			q.attrVals = append(q.attrVals, v)
		}
	} else if q.mode == modeAdd {
		for k, v := range vals {
			q.attrs = append(q.attrs, k)
			q.attrVals = append(q.attrVals, v)
		}
	}
	return q
}

func (q *Query) addSelection(key string, val interface{}, op string) {
	q.selection = append(q.selection, key+op)
	q.selectionVals = append(q.selectionVals, val)
}

func (q *Query) build() (query string) {
	if q.table == "" {
		return
	}
	switch q.mode {
	case modeGet:
		query = q.buildGet()
	case modeAdd:
		query = q.buildAdd()
	case modeUpdate:
		query = q.buildUpdate()
	case modeDelete:
		query = q.buildDelete()
	}
	return
}

func (q *Query) buildAdd() string {
	if len(q.attrs) < 1 || len(q.attrs) != len(q.attrVals) {
		return ""
	}
	vs := make([]string, len(q.attrs))
	for i := 0; i < len(vs); i++ {
		vs[i] = q.driver.Param()
	}
	return fmt.Sprintf("insert into %s (%s) values (%s)",
		q.table, strings.Join(q.attrs, ","), strings.Join(vs, ","))
}

func (q *Query) buildDelete() string {
	if len(q.selection) < 1 {
		return ""
	}
	sel := strings.Join(q.selection, " ")
	return fmt.Sprintf("delete from %s where %s", q.table, sel)
}

func (q *Query) buildGet() string {
	var ord string
	if q.order != "" && q.orderAttr != "" {
		ord = "order by " + q.orderAttr + " collate nocase " + q.order
	}
	if len(q.selection) < 1 {
		q.All()
	}
	if len(q.attrs) < 1 {
		q.attrs = append(q.attrs, "*")
	}
	pro := strings.Join(q.attrs, ",")
	sel := strings.Join(q.selection, " ")
	return fmt.Sprintf("select %s from %s where %s %s", pro, q.table, sel, ord)
}

func (q *Query) buildUpdate() string {
	if len(q.attrs) < 1 || len(q.selection) < 1 {
		return ""
	}
	attrs := strings.Join(q.attrs, ",")
	sel := strings.Join(q.selection, " ")
	return fmt.Sprintf("update %s set %s where %s", q.table, attrs, sel)
}
