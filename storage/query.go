package storage

import (
	"fmt"
	"sail/conf"
	"sail/errors"
	"strings"
)

const (
	and  = " and "
	or   = " or "
	asc  = " asc "
	desc = " desc "
)

const (
	modeGet    = 0x001
	modeAdd    = 0x002
	modeUpdate = 0x004
	modeDelete = 0x008
)

type Query struct {
	mode          uint8
	table         string
	attrs         []string
	attrVals      []interface{}
	selection     []string
	selectionVals []interface{}
	order         string
	orderAttr     string
}

func Add() *Query {
	return &Query{mode: modeAdd}
}

func Delete() *Query {
	return &Query{mode: modeDelete}
}

func Get() *Query {
	return &Query{mode: modeGet}
}

func Update() *Query {
	return &Query{mode: modeUpdate}
}

func (q *Query) All() *Query {
	q.selection = append(q.selection, "1=1")
	return q
}

func (q *Query) And() *Query {
	if len(q.selection) > 1 {
		q.selection = append(q.selection, and)
	}
	return q
}

func (q *Query) Asc() *Query {
	q.order = asc
	return q
}

func (q *Query) Attrs(attrs ...string) *Query {
	q.attrs = attrs
	return q
}

func (q *Query) Desc() *Query {
	q.order = desc
	return q
}

func (q *Query) Equals(key string, val interface{}) *Query {
	q.addSelection(key, val, "=?")
	return q
}

func (q *Query) Exec() (res interface{}) {
	var err error
	switch q.mode {
	case modeGet:
		res, err = DB().Query(q.build(), q.selectionVals...) // res = sql.Rows
	case modeAdd, modeUpdate, modeDelete:
		vals := append(q.attrVals, q.selectionVals...)
		_, err = DB().Exec(q.build(), vals...)
		res = (err == nil) // res = bool
	}
	if err != nil {
		errors.Log(err, conf.Instance().DevMode)
	}
	return
}

func (q *Query) In(table string) *Query {
	q.table = table
	return q
}

func (q *Query) Or() *Query {
	if len(q.selection) > 1 {
		q.selection = append(q.selection, or)
	}
	return q
}

func (q *Query) Order(attr string) *Query {
	q.orderAttr = attr
	return q
}

func (q *Query) String() string {
	switch q.mode {
	case modeGet:
		a := strings.Join(q.attrs, ", ")
		s := strings.Join(q.selection, " ")
		o := ""
		if q.order != "" && q.orderAttr != "" {
			o = "order by " + q.orderAttr + q.order
		}
		for i := 0; strings.Contains(s, "?"); i++ {
			s = strings.Replace(s, "?", fmt.Sprintf("%v", q.selectionVals[i]), 1)
		}
		return fmt.Sprintf("select %s from %s where %s %s",
			a, q.table, s, o)
	case modeAdd:
		a := strings.Join(q.attrs, ", ")
		v := ""
		for _, i := range q.attrVals {
			v += fmt.Sprintf("%v, ", i)
		}
		return fmt.Sprintf("insert into %s (%s) values (%s)",
			q.table, a, v[:len(v)-2])
	case modeUpdate:
		a := strings.Join(q.attrs, ", ")
		s := strings.Join(q.selection, ", ")
		for i := 0; strings.Contains(a, "?"); i++ {
			a = strings.Replace(a, "?", fmt.Sprintf("%v", q.attrVals[i]), 1)
		}
		for i := 0; strings.Contains(s, "?"); i++ {
			s = strings.Replace(s, "?", fmt.Sprintf("%v", q.selectionVals[i]), 1)
		}
		return fmt.Sprintf("update %s set %s where %s", q.table, a, s)
	case modeDelete:
		s := ""
		if len(q.selection) > 0 {
			s = "where " + strings.Join(q.selection, " ")
		}
		for i := 0; strings.Contains(s, "?"); i++ {
			s = strings.Replace(s, "?", fmt.Sprintf("%v", q.selectionVals[i]), 1)
		}
		return fmt.Sprintf("delete from %s %s", q.table, s)
	default:
		return ""
	}
}

func (q *Query) Values(vals map[string]interface{}) *Query {
	if q.mode == modeUpdate {
		for k, v := range vals {
			q.attrs = append(q.attrs, k+"=?")
			q.attrVals = append(q.attrVals, v)
		}
	} else {
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
		vs[i] = "?"
	}
	return fmt.Sprintf("insert into %s (%s) values (%s)",
		q.table, strings.Join(q.attrs, ","), strings.Join(vs, ","))
}

func (q *Query) buildDelete() string {
	if len(q.selection) < 1 {
		// what happens when there is no selection attribute?
		// For now, we make sure that nothing gets deleted. This
		// might change in the future, allowing users to delete
		// all accounts at once.
		q.selection = append(q.selection, "1=0")
	}
	sel := strings.Join(q.selection, " ")
	return fmt.Sprintf("delete from %s where %s", q.table, sel)
}

func (q *Query) buildGet() string {
	var ord string
	if q.order != "" && q.orderAttr != "" {
		ord = "order by " + q.orderAttr + q.order
	}
	if len(q.selection) < 1 {
		q.All()
	}
	pro := strings.Join(q.attrs, ",")
	sel := strings.Join(q.selection, " ")
	return fmt.Sprintf("select %s from %s where %s %s", pro, q.table, sel, ord)
}

func (q *Query) buildUpdate() string {
	attrs := strings.Join(q.attrs, ",")
	sel := strings.Join(q.selection, " ")
	return fmt.Sprintf("update %s set %s where %s", q.table, attrs, sel)
}
