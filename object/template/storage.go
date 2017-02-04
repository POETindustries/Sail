package template

import (
	"database/sql"
	"sail/conf"
	"sail/errors"
	"sail/object/schema"
	"sail/store"
)

func fromStorageByID(ids ...uint32) []*Template {
	query := store.Get().In("sl_template").Attrs(schema.TemplateAttrs...)
	if len(ids) == 1 {
		query.Equals(schema.TemplateID, ids[0])
	} else if len(ids) > 1 {
		// query.EqualsMany(schema.TemplateID, ids)
	}
	rows, _ := query.Exec()
	ts := scanTemplate(rows)
	for _, t := range ts {
		rows, _ = store.Get().In("sl_template_widgets").
			Equals(schema.TemplateID, t.ID).
			Attrs(schema.TemplateWidgetID).Exec()
		t.WidgetIDs = scanWidgetID(rows)
	}
	return ts
}

func scanTemplate(rows *sql.Rows) []*Template {
	defer rows.Close()
	var ts []*Template
	for rows.Next() {
		t := New()
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			errors.Log(err, conf.Instance().DevMode)
			return nil
		}
		ts = append(ts, t)
	}
	return ts
}

func scanWidgetID(rows *sql.Rows) []uint32 {
	defer rows.Close()
	var ws []uint32
	for rows.Next() {
		var w uint32
		if err := rows.Scan(&w); err != nil {
			errors.Log(err, conf.Instance().DevMode)
			return nil
		}
		ws = append(ws, w)
	}
	return ws
}
