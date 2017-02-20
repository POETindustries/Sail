package widget

import (
	"database/sql"
	"fmt"
	"sail/log"
	"sail/object/schema"
	"sail/store"
)

func fromStorageByID(ids ...uint32) []*Widget {
	query := store.Get().In("sl_widget").Attrs(schema.WidgetAttrs...)
	if len(ids) == 1 {
		query.Equals(schema.WidgetID, ids[0])
	} else if len(ids) > 1 {
		// query.EqualsMany(schema.WidgetID, ids)
	}
	rows, _ := query.Exec()
	return scanWidget(rows)
}

func fetchData(widget *Widget) {
	switch widget.Type {
	case "nav":
		widget.Data = fetchNavData(widget.ID)
	case "text":
		widget.Data = fetchTextData(widget.ID)
	}
}

func fetchNavData(id uint32) *Nav {
	rows, _ := store.Get().In("sl_widget_nav").Attrs(schema.NavAttrs...).
		Equals(schema.NavWidgetID, id).Exec()
	return scanNav(rows)
}

func fetchTextData(id uint32) *Text {
	rows, _ := store.Get().In("sl_widget_text").Equals(schema.TextID, id).Exec()
	ts := scanText(rows)
	if len(ts) < 1 {
		// could return &Text{}, but nil is consistent with other
		// funcs. It needs to be seen whether this is a good idea.
		return nil
	}
	return ts[0]
}

func scanWidget(rows *sql.Rows) []*Widget {
	defer rows.Close()
	var ws []*Widget
	for rows.Next() {
		w := New()
		if err := rows.Scan(&w.ID, &w.Name, &w.RefName, &w.Type); err != nil {
			log.DB(err, log.LvlWarn)
			return nil
		}
		ws = append(ws, w)
	}
	return ws
}

func scanNav(rows *sql.Rows) *Nav {
	defer rows.Close()
	n := Nav{}
	for rows.Next() {
		e := NavEntry{}
		if err := rows.Scan(&e.ID, &e.Name, &e.RefID, &e.Submenu, &e.Pos); err != nil {
			log.DB(err, log.LvlWarn)
			return nil
		}
		e.RefURL = fmt.Sprintf("uuid/%d", e.RefID)
		n.Entries = append(n.Entries, &e)
	}
	return &n
}

func scanText(rows *sql.Rows) []*Text {
	defer rows.Close()
	var ts []*Text
	for rows.Next() {
		t := Text{}
		if err := rows.Scan(&t.Content); err != nil {
			log.DB(err, log.LvlWarn)
			return nil
		}
		ts = append(ts, &t)
	}
	return ts
}
