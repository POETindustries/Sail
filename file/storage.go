package file

import (
	"database/sql"
	"sail/conf"
	"sail/errors"
	"sail/object/schema"
	"sail/storage"
)

func fromStorageChildren(id uint32) []*File {
	rows := storage.Get().In("sl_object").Attrs(schema.ObjectAttrs...).
		Equals(schema.ObjectParent, id).And().
		Equals(schema.ObjectTypeMajor, Text).And().
		Equals(schema.ObjectTypeMinor, Html).Exec()
	return scanChildren(rows.(*sql.Rows))
}

func scanChildren(rows *sql.Rows) []*File {
	defer rows.Close()
	var fs []*File
	for rows.Next() {
		f := File{}
		if err := rows.Scan(&f.ID, &f.Name, &f.machineName, &f.parent,
			&f.mimeTypeMajor, &f.mimeTypeMinor, &f.status, &f.owner,
			&f.cDate, &f.eDate, &f.Address); err != nil {
			errors.Log(err, conf.Instance().DevMode)
			return nil
		}
		fs = append(fs, &f)
	}
	return fs
}
