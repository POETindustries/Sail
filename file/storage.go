package file

import (
	"database/sql"
	"sail/conf"
	"sail/errors"
	"sail/file/schema"
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
