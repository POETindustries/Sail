package template

import (
	"database/sql"
	"sail/conf"
	"sail/errors"
	"sail/page/schema"
	"sail/storage"
)

func fromStorageByID(ids ...uint32) []*Template {
	query := storage.Get().In("sl_template").Attrs(schema.TemplateAttrs...)
	if len(ids) == 1 {
		query.Equals(schema.TemplateID, ids[0])
	} else if len(ids) > 1 {
		// query.EqualsMany(schema.TemplateID, ids)
	}
	rows := query.Exec()
	ts := scanTemplate(rows.(*sql.Rows))
	for _, t := range ts {
		rows = storage.Get().In("sl_template_widgets").
			Equals(schema.TemplateID, t.ID).
			Attrs(schema.TemplateWidgetID).Exec()
		t.WidgetIDs = scanWidgetID(rows.(*sql.Rows))
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
