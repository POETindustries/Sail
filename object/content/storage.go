package content

import (
	"database/sql"
	"sail/conf"
	"sail/errors"
	"sail/object/schema"
	"sail/storage"
)

func fromStorageByURL(url string) []*Content {
	t := "sl_content natural join sl_meta"
	a := append(schema.ContentAttrs, schema.MetaAttrs...)
	rows := storage.Get().In(t).Attrs(a...).Equals(schema.ContentURL, url).Exec()
	return scan(rows.(*sql.Rows))
}

func fromStorageByID(id uint32) []*Content {
	t := "sl_content natural join sl_meta"
	a := append(schema.ContentAttrs, schema.MetaAttrs...)
	rows := storage.Get().In(t).Attrs(a...).Equals(schema.ContentID, id).Exec()
	return scan(rows.(*sql.Rows))
}

func scan(rows *sql.Rows) []*Content {
	defer rows.Close()
	var cs []*Content
	for rows.Next() {
		c := New()
		if err := rows.Scan(&c.ID, &c.Title, &c.Content, &c.Meta.ID,
			&c.TemplateID, &c.URL, &c.Parent, &c.Status, &c.Owner, &c.CDate,
			&c.EDate, &c.Meta.Title, &c.Meta.Keywords, &c.Meta.Description,
			&c.Meta.Language, &c.Meta.PageTopic, &c.Meta.RevisitAfter,
			&c.Meta.Robots); err != nil {
			errors.Log(err, conf.Instance().DevMode)
			return nil
		}
		cs = append(cs, c)
	}
	return cs
}
