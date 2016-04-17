package content

import (
	"database/sql"
	"sail/conf"
	"sail/errors"
	"sail/object/schema"
	"sail/storage"
)

const queryMinByURL = `select * from ((select ` + schema.ObjectID + `,
	` + schema.ObjectName + `,` + schema.ObjectURLCache + ` from sl_object
	where ` + schema.ObjectStatus + `=1 and ` + schema.ObjectURLCache + `=?)
	natural join sl_content) natural join sl_meta;`

const queryMinByID = `select * from ((select ` + schema.ObjectID + `,
	` + schema.ObjectName + `,` + schema.ObjectURLCache + ` from sl_object
	where ` + schema.ObjectStatus + `=1 and ` + schema.ObjectID + `=?)
	natural join sl_content) natural join sl_meta;`

func fromStorageMinByURL(url string) *Content {
	row := storage.DB().QueryRow(queryMinByURL, url)
	return scan(row)
}

func fromStorageMinByID(id uint32) *Content {
	row := storage.DB().QueryRow(queryMinByID, id)
	return scan(row)
}

func fromStorageFullByURL(url string) []*Content {
	t := "(sl_object natural join sl_content) natural join sl_meta"
	a := append(schema.ObjectAttrs, schema.ContentContent, schema.ContentMetaID,
		schema.ContentTemplateID)
	a = append(a, schema.MetaAttrs...)
	rows := storage.Get().In(t).Attrs(a...).Equals(schema.ObjectURLCache, url).Exec()
	return scanFull(rows.(*sql.Rows))
}

func fromStorageFullByID(id uint32) []*Content {
	t := "(sl_object natural join sl_content) natural join sl_meta"
	a := append(schema.ContentAttrs, schema.MetaAttrs...)
	rows := storage.Get().In(t).Attrs(a...).Equals(schema.ContentID, id).Exec()
	return scanFull(rows.(*sql.Rows))
}

func scan(row *sql.Row) *Content {
	c := New()
	if err := row.Scan(&c.ID, &c.Title, &c.URL, &c.Content,
		&c.Meta.ID, &c.TemplateID, &c.Meta.Title, &c.Meta.Keywords,
		&c.Meta.Description, &c.Meta.Language, &c.Meta.PageTopic,
		&c.Meta.RevisitAfter, &c.Meta.Robots); err != nil {
		errors.Log(err, conf.Instance().DevMode)
		return nil
	}
	return c
}

func scanFull(rows *sql.Rows) []*Content {
	defer rows.Close()
	var cs []*Content
	for rows.Next() {
		c := New()
		var maj, min uint16
		if err := rows.Scan(&c.ID, &c.Title, &c.MachineName, &c.Parent, &maj,
			&min, &c.Status, &c.Owner, &c.CDate, &c.EDate, &c.URL, &c.Content,
			&c.Meta.ID, &c.TemplateID, &c.Meta.Title, &c.Meta.Keywords,
			&c.Meta.Description, &c.Meta.Language, &c.Meta.PageTopic,
			&c.Meta.RevisitAfter, &c.Meta.Robots); err != nil {
			errors.Log(err, conf.Instance().DevMode)
			return nil
		}
		cs = append(cs, c)
	}
	return cs
}
