package object

import (
	"sail/object/schema"
	"sail/storage"
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
	if rows, err := storage.DB().Query(queryStaticAddr, id); err == nil {
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
