package user

import (
	"database/sql"
	"sail/conf"
	"sail/errors"
	"sail/storage"
	"sail/user/schema"
)

func fromStorageByName(names ...string) []*User {
	query := storage.Get().In("sl_user").Attrs(schema.UserAttrs...)
	if len(names) == 1 {
		query.Equals(schema.UserName, names[0])
	} else if len(names) > 1 {
		// query.EqualsMany(schema.UserName, names)
	}
	rows := query.Exec()
	return scanUser(rows.(*sql.Rows))
}

func scanUser(rows *sql.Rows) []*User {
	defer rows.Close()
	var us []*User
	for rows.Next() {
		u := User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.pass, &u.FirstName, &u.LastName,
			&u.Email, &u.Phone, &u.CDate, &u.ExpDate); err != nil {
			errors.Log(err, conf.Instance().DevMode)
			return nil
		}
		us = append(us, &u)
	}
	return us
}
