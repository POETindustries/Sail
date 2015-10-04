package widgetstore

import (
	"database/sql"
	"fmt"
	"sail/storage/psqldb"
	"sail/widget"
)

// Query collects all information needed for querying the database.
type Query struct {
	query *psqldb.Query
}

// ByID prepares the query to select the widget that matches the given id.
func (q *Query) ByID(ids ...uint32) *Query {
	for _, id := range ids {
		q.query.AddAttr(widgetID, id, psqldb.OpOr)
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
func (q *Query) Widgets() ([]*widget.Widget, error) {
	q.query.Table = "sl_widget"
	q.query.Proj = widgetAttrs
	return q.scanWidgets(q.query.Execute())
}

// Menu executes the query, collecting information for one menu widget.
func (q *Query) Menu() (*widget.Menu, error) {
	stmt := "sl_widget_menu join (select %s,%s from sl_page) as p on %s=%s"
	q.query.Table = fmt.Sprintf(stmt, expPageID, expPageURL, expPageID, menuEntryReferenceID)
	q.query.Proj = menuAttrs + "," + expPageURL

	if o := q.query.Order(); o != "" {
		q.query.SetOrderStmt("order by " + menuEntryPosition + o)
	}
	return q.scanMenu(q.query.Execute())
}

// TextField executes the query, providing a text widget in return.
func (q *Query) TextField() (*widget.Text, error) {
	q.query.Table = "sl_widget_text"
	q.query.Proj = textContent
	return q.scanTextField(q.query.Execute())
}

func (q *Query) scanWidgets(data *sql.Rows, err error) ([]*widget.Widget, error) {
	if err != nil {
		return nil, err
	}
	widgets := []*widget.Widget{}
	defer data.Close()
	for data.Next() {
		w := widget.Widget{}
		if err = data.Scan(&w.ID, &w.Name, &w.RefName, &w.Type); err != nil {
			return nil, err
		}
		widgets = append(widgets, &w)
	}
	return widgets, nil
}

func (q *Query) scanMenu(data *sql.Rows, err error) (*widget.Menu, error) {
	if err != nil {
		return nil, err
	}
	menu := widget.Menu{}
	defer data.Close()
	for data.Next() {
		e := widget.MenuEntry{}
		if err = data.Scan(&e.ID, &e.Name, &e.RefID, &e.Submenu, &e.Pos, &e.RefURL); err != nil {
			return nil, err
		}
		menu.Entries = append(menu.Entries, &e)
	}
	return &menu, nil
}

func (q *Query) scanTextField(data *sql.Rows, err error) (*widget.Text, error) {
	if err != nil {
		return nil, err
	}
	text := widget.Text{}
	defer data.Close()
	for data.Next() {
		if err = data.Scan(&text.Content); err != nil {
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
