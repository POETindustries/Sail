package widget

import (
	"database/sql"
	"fmt"
	"sail/conf"
	"sail/errors"
	"sail/object/schema"
	"sail/storage"
)

func fromStorageByID(ids ...uint32) []*Widget {
	query := storage.Get().In("sl_widget").Attrs(schema.WidgetAttrs...)
	if len(ids) == 1 {
		query.Equals(schema.WidgetID, ids[0])
	} else if len(ids) > 1 {
		// query.EqualsMany(schema.WidgetID, ids)
	}
	rows := query.Exec()
	return scanWidget(rows.(*sql.Rows))
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
	stmt := "sl_widget_nav natural join (select %s,%s from sl_content)"
	t := fmt.Sprintf(stmt, schema.ContentID, schema.ContentURL)
	a := append(schema.NavAttrs, schema.ContentURL)
	rows := storage.Get().In(t).Attrs(a...).
		Equals(schema.NavWidgetID, id).Exec()
	return scanNav(rows.(*sql.Rows))
}

func fetchTextData(id uint32) *Text {
	rows := storage.Get().In("sl_widget_text").Equals(schema.TextID, id).Exec()
	ts := scanText(rows.(*sql.Rows))
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
			errors.Log(err, conf.Instance().DevMode)
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
		if err := rows.Scan(&e.ID, &e.Name, &e.RefID, &e.Submenu, &e.Pos,
			&e.RefURL); err != nil {
			errors.Log(err, conf.Instance().DevMode)
			return nil
		}
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
			errors.Log(err, conf.Instance().DevMode)
			return nil
		}
		ts = append(ts, &t)
	}
	return ts
}
