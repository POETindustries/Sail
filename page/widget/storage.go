package widget

import (
	"database/sql"
	"fmt"
	"sail/page/data"
	"sail/storage/psqldb"
	"sail/storage/schema"
)

// Query collects all information needed for querying the database.
type Query struct {
	query *psqldb.Query
}

// ByID prepares the query to select the widget that matches the given id.
func (q *Query) ByID(ids ...uint32) *Query {
	for _, id := range ids {
		q.query.AddAttr(schema.WidgetID, id, psqldb.OpOr)
	}
	return q
}

// Ascending prepares the query for table-specific ordering.
func (q *Query) Ascending() *Query {
	q.query.Ascending()
	return q
}

// Descending prepares the query for table-specific ordering.
func (q *Query) Descending() *Query {
	q.query.Descending()
	return q
}

// Widgets executes the query and returns all matching widget objects.
func (q *Query) Widgets() ([]*data.Widget, error) {
	q.query.Table = "sl_widget"
	q.query.Proj = schema.WidgetAttrs
	return q.scanWidgets(q.query.Execute())
}

// Menu executes the query, collecting information for one menu data.
func (q *Query) Menu() (*data.Menu, error) {
	stmt := "sl_widget_menu join (select %s,%s from sl_page) as p on %s=%s"
	q.query.Table = fmt.Sprintf(stmt,
		schema.PageID, schema.PageURL, schema.PageID, schema.MenuEntryRefID)
	q.query.Proj = schema.MenuAttrs + "," + schema.PageURL

	if o := q.query.Order(); o != "" {
		q.query.SetOrderStmt("order by " + schema.MenuEntryPosition + o)
	}
	return q.scanMenu(q.query.Execute())
}

// TextField executes the query, providing a text widget in return.
func (q *Query) TextField() (*data.Text, error) {
	q.query.Table = "sl_widget_text"
	q.query.Proj = schema.TextContent
	return q.scanTextField(q.query.Execute())
}

func (q *Query) scanWidgets(rows *sql.Rows, err error) ([]*data.Widget, error) {
	if err != nil {
		return nil, err
	}
	widgets := []*data.Widget{}
	defer rows.Close()
	for rows.Next() {
		w := data.NewWidget()
		if err = rows.Scan(&w.ID, &w.Name, &w.RefName, &w.Type); err != nil {
			return nil, err
		}
		widgets = append(widgets, w)
	}
	return widgets, nil
}

func (q *Query) scanMenu(rows *sql.Rows, err error) (*data.Menu, error) {
	if err != nil {
		return nil, err
	}
	menu := data.Menu{}
	defer rows.Close()
	for rows.Next() {
		e := data.MenuEntry{}
		if err = rows.Scan(&e.ID, &e.Name, &e.RefID, &e.Submenu, &e.Pos, &e.RefURL); err != nil {
			return nil, err
		}
		menu.Entries = append(menu.Entries, &e)
	}
	return &menu, nil
}

func (q *Query) scanTextField(rows *sql.Rows, err error) (*data.Text, error) {
	if err != nil {
		return nil, err
	}
	text := data.Text{}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&text.Content); err != nil {
			return nil, err
		}
	}
	return &text, nil
}

// Get starts building the query that gets sent to the database.
//
// TODO: describe how queries should be built using method chaining.
func Get() *Query {
	return &Query{query: &psqldb.Query{}}
}
