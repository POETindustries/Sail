package group

import (
	"database/sql"
	"sail/log"
	"sail/store"
	"sail/user/rights"
	"sail/user/schema"
)

func fromStorageByID(id ...uint32) []*Group {
	query := store.Get().In("sl_group").Attrs(schema.GroupAttrs...)
	if len(id) == 1 {
		query.Equals(schema.GroupID, id[0])
	} else if len(id) > 1 {
		// query.EqualsMany(schema.GroupID, id)
	}
	rows, _ := query.Exec()
	gs := scanGroup(rows)
	for _, g := range gs {
		g.users = fetchMembers(g.ID)
	}
	return gs
}

func fetchMembers(id uint32) map[uint32]bool {
	rows, _ := store.Get().In("sl_group_members").Attrs(schema.UserID).
		Equals(schema.GroupID, id).Exec()
	return scanMembers(rows)
}

func scanGroup(rows *sql.Rows) []*Group {
	defer rows.Close()
	var gs []*Group
	for rows.Next() {
		g := New()
		if err := rows.Scan(&g.ID, &g.Name, &g.perm[rights.Maintenance],
			&g.perm[rights.Users], &g.perm[rights.Content],
			&g.perm[rights.Config]); err != nil {
			log.DB(err, log.LvlWarn)
		}
		gs = append(gs, g)
	}
	return gs
}

func scanMembers(rows *sql.Rows) map[uint32]bool {
	defer rows.Close()
	ms := make(map[uint32]bool)
	for rows.Next() {
		var id uint32
		if err := rows.Scan(&id); err != nil {
			log.DB(err, log.LvlWarn)
		}
		ms[id] = true
	}
	return ms
}
