package templatestore

import (
	"database/sql"
	"sail/storage/psqldb"
	"sail/storage/schema"
	"sail/tmpl"
)

// Query collects all information needed for querying the database.
type Query struct {
	query *psqldb.Query
}

// ByID prepares the query to select the page that matches the given id.
func (q *Query) ByID(ids ...uint32) *Query {
	for _, id := range ids {
		q.query.AddAttr(schema.TemplateID, id, psqldb.OpOr)
	}
	return q
}

// Templates executes the query and returns all matching widget objects.
func (q *Query) Templates() ([]*tmpl.Template, error) {
	q.query.Table = "sl_template"
	q.query.Proj = schema.TemplateAttrs
	return q.scanTemplates(q.query.Execute())
}

// WidgetIDs executes the query and returns the ids of all widgets
// used in this template.
func (q *Query) WidgetIDs() ([]uint32, error) {
	q.query.Table = "sl_template_widgets"
	q.query.Proj = schema.WidgetID
	return q.scanWidgetIDs(q.query.Execute())
}

func (q *Query) scanTemplates(data *sql.Rows, err error) ([]*tmpl.Template, error) {
	if err != nil {
		return nil, err
	}
	ts := []*tmpl.Template{}
	defer data.Close()
	for data.Next() {
		t := tmpl.New()
		if err = data.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}

func (q *Query) scanWidgetIDs(data *sql.Rows, err error) ([]uint32, error) {
	if err != nil {
		return nil, err
	}
	var ids []uint32
	defer data.Close()
	for data.Next() {
		var id uint32
		if err = data.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// Get starts building the query that gets sent to the database.
//
// TODO: describe how queries should be built using method chaining.
func Get() *Query {
	return &Query{query: &psqldb.Query{}}
}
