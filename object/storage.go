package object

import (
	"sail/conf"
	"sail/errors"
	"sail/object/schema"
	"sail/store"
	"strings"
)

const queryStaticAddr = `with recursive parent_of(
` + schema.ObjectID + `,
` + schema.ObjectParent + `,
` + schema.ObjectMachineName + `,
` + schema.ObjectStatus + `)
as (
    select ` + schema.ObjectID + `,
	` + schema.ObjectParent + `,
	` + schema.ObjectMachineName + `,
    ` + schema.ObjectStatus + `
	from sl_object where ` + schema.ObjectID + `=?
    union
    select sl_object.` + schema.ObjectID + `,
	sl_object.` + schema.ObjectParent + `,
	sl_object.` + schema.ObjectMachineName + `,
    sl_object.` + schema.ObjectStatus + `
	from sl_object,parent_of
	where parent_of.` + schema.ObjectParent + `=sl_object.` + schema.ObjectID + `)
select ` + schema.ObjectMachineName + `,` + schema.ObjectStatus + ` from parent_of;`

func fromStorageStaticAddr(id uint32, public bool) (addr string) {
	rows, _ := store.Get().In("sl_object").Attrs(schema.ObjectURLCache).
		Equals(schema.ObjectID, id).Exec()
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&addr); err != nil {
		return fromStorageBuildStaticAddr(id, public)
	}
	return
}

func fromStorageBuildStaticAddr(id uint32, public bool) (addr string) {
	if rows, err := store.DB().Query(queryStaticAddr, id); err == nil {
		defer rows.Close()
		var p int8
		for rows.Next() {
			var n string
			if err = rows.Scan(&n, &p); err != nil || (public && (p != 1)) {
				return ""
			}
			addr = "/" + n + addr
		}
	}
	return
}

func fromStorageID(url string, public bool) (id uint32) {
	rows, _ := store.Get().In("sl_object").Attrs(schema.ObjectID).
		Equals(schema.ObjectURLCache, url).Exec()
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&id); err != nil {
		return fromStorageBuildID(url, public)
	}
	return
}

func fromStorageBuildID(url string, public bool) uint32 {
	locs := strings.Split(url, "/")
	rows, _ := store.Get().In("sl_object").Attrs(schema.ObjectID, schema.ObjectStatus).
		Equals(schema.ObjectMachineName, locs[len(locs)-1]).Exec()
	var ids []uint32
	for rows.Next() {
		var id uint32
		var p int8
		if err := rows.Scan(&id, &p); err != nil || (public && p != 1) {
			errors.Log(err, conf.Instance().DevMode)
			continue
		}
		ids = append(ids, id)
	}
	for _, i := range ids {
		if fromStorageStaticAddr(i, public) == url {
			return i
		}
	}
	return 0
}
