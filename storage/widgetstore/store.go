package widgetstore

import (
	"bytes"
	"database/sql"
	"fmt"
	"sail/errors"
	"sail/storage/conn"
	"sail/widget"
	"strconv"
)

const and, or = " and ", " or "

// Query collects all information needed for querying the database.
type Query struct {
	table     string
	proj      string
	sel       bytes.Buffer
	selAttrs  []interface{}
	attrCount int
	order     string
	asc       bool
	desc      bool
}

// ByID prepares the query to select the widget that matches the given id.
func (q *Query) ByID(ids ...uint32) *Query {
	for _, id := range ids {
		q.addAttr(widgetID, id, or)
	}
	return q
}

// Ascending prepares the query for table-specific ordering.
func (q *Query) Ascending() *Query {
	q.asc = true
	return q
}

// Descending prepares the query for table-specific ordering.
func (q *Query) Descending() *Query {
	q.desc = true
	return q
}

// Widgets executes the query and returns all matching widget objects.
func (q *Query) Widgets() ([]*widget.Widget, error) {
	q.table = "sl_widget"
	q.proj = widgetAttrs
	return q.scanWidgets(q.execute())
}

// Menu executes the query, collecting information for one menu widget.
func (q *Query) Menu() (*widget.Menu, error) {
	stmt := "sl_widget_menu join (select %s,%s from sl_page) as p on %s=%s"
	q.table = fmt.Sprintf(stmt, expPageID, expPageURL, expPageID, menuEntryReferenceID)
	q.proj = menuAttrs + "," + expPageURL

	if q.asc {
		q.order = "order by " + menuEntryPosition + " asc"
	} else if q.desc {
		q.order = "order by " + menuEntryPosition + " desc"
	}
	return q.scanMenu(q.execute())
}

// TextField executes the query, providing a text widget in return.
func (q *Query) TextField() (*widget.Text, error) {
	q.table = "sl_widget_text"
	q.proj = textContent
	return q.scanTextField(q.execute())
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

func (q *Query) addAttr(key string, val interface{}, op string) {
	if q.attrCount > 1 {
		q.sel.WriteString(op)
	}
	q.sel.WriteString(key + "=$" + strconv.Itoa(q.attrCount))
	q.selAttrs = append(q.selAttrs, val)
	q.attrCount++
}

func (q *Query) execute() (*sql.Rows, error) {
	if q.sel.Len() < 1 {
		return nil, errors.NoArguments()
	}
	return conn.Instance().DB.Query(q.build(), q.selAttrs...)
}

func (q *Query) build() string {
	query := "select %s from %s where %s %s"
	return fmt.Sprintf(query, q.proj, q.table, q.sel.String(), q.order)
}

// Get starts building the query that gets sent to the database.
//
// TODO: describe how queries should be built using method chaining.
func Get() *Query {
	return &Query{attrCount: 1}
}
