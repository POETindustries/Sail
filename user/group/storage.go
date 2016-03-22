package group

import (
	"database/sql"
	"sail/conf"
	"sail/errors"
	"sail/storage"
	"sail/user/permission"
	"sail/user/schema"
)

func fromStorageByID(id ...uint32) []*Group {
	query := storage.Get().In("sl_group").Attrs(schema.GroupAttrs...)
	if len(id) == 1 {
		query.Equals(schema.GroupID, id[0])
	} else if len(id) > 1 {
		// query.EqualsMany(schema.GroupID, id)
	}
	rows := query.Exec()
	return scanGroup(rows.(*sql.Rows))
}

func scanGroup(rows *sql.Rows) []*Group {
	defer rows.Close()
	var gs []*Group
	for rows.Next() {
		g := Group{}
		if err := rows.Scan(&g.ID, &g.Name, &g.perm[permission.Maintenance],
			&g.perm[permission.Users]); err != nil {
			errors.Log(err, conf.Instance().DevMode)
		}
		gs = append(gs, &g)
	}
	return gs
}
