package file

import (
	"database/sql"
	"sail/log"
	"sail/object/schema"
	"sail/store"
)

func fromStorageChildren(id uint32, includeCurrent bool) []*File {
	query := store.Get().In("sl_object").Attrs(schema.ObjectAttrs...).
		Equals(schema.ObjectParent, id)
	if includeCurrent {
		query.Or().Equals(schema.ObjectID, id).And().
			NotEquals(schema.ObjectTypeMajor, Directory)
	}
	rows, _ := query.Order(schema.ObjectName).Asc().Exec()
	return scanChildren(rows)
}

func fromStorageChildCount(id uint32) (count uint32) {
	rows, _ := store.Get().In("sl_object").Attrs("count("+schema.ObjectParent+")").
		Equals(schema.ObjectParent, id).Exec()
	defer rows.Close()
	rows.Next()
	rows.Scan(&count)
	return
}

func scanChildren(rows *sql.Rows) []*File {
	defer rows.Close()
	var fs []*File
	for rows.Next() {
		f := File{}
		if err := rows.Scan(&f.ID, &f.Name, &f.machineName, &f.parent,
			&f.mimeTypeMajor, &f.mimeTypeMinor, &f.status, &f.owner,
			&f.cDate, &f.eDate, &f.Address); err != nil {
			log.DB(err, log.LvlWarn)
			return nil
		}
		fs = append(fs, &f)
	}
	return fs
}
