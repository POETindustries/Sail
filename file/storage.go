package file

import (
	"database/sql"
	"sail/conf"
	"sail/errors"
	"sail/file/schema"
	cSchema "sail/object/schema"
	"sail/storage"
)

func fromStorageGetAddr(uuid string, public bool) (addr string) {
	query := storage.Get().In("sl_file").Attrs(schema.FileAddr).
		Equals(schema.FileID, uuid[5:])
	if public {
		query.And().Equals(schema.FileStatus, Public)
	}
	rows := query.Exec().(*sql.Rows)
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&addr); err != nil {
			errors.Log(err, conf.Instance().DevMode)
		}
	}
	return
}

func fromStorageAsContent(dir string) []*File {
	rows := storage.Get().In("sl_content").
		Attrs(cSchema.ContentTitle, cSchema.ContentURL, cSchema.ContentStatus).
		Equals(cSchema.ContentParent, dir).Or().Equals(cSchema.ContentURL, dir).
		Exec()
	return scanContent(rows.(*sql.Rows))
}

func scanContent(rows *sql.Rows) []*File {
	defer rows.Close()
	var fs []*File
	for rows.Next() {
		f := File{mimeType: 1}
		if err := rows.Scan(&f.Name, &f.Address, &f.status); err != nil {
			errors.Log(err, conf.Instance().DevMode)
			return nil
		}
		fs = append(fs, &f)
	}
	return fs
}
