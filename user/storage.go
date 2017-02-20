package user

import (
	"database/sql"
	"sail/log"
	"sail/store"
	"time"
)

const (
	userTable     = "sl_usr"
	userID        = "usr_id"
	userName      = "usr_name"
	userPass      = "usr_pass"
	userFirstName = "usr_firstname"
	userLastName  = "usr_lastname"
	userEmail     = "usr_email"
	userPhone     = "usr_phone"
	userCDate     = "usr_cdate"
	userExpDate   = "usr_expdate"
)

var userAttrs = []string{userID, userName, userPass, userFirstName,
	userLastName, userEmail, userPhone, userCDate, userExpDate}

// SetupData contains data needed to setup persistent storage.
func SetupData() (string, []*store.SetupData) {
	return userTable, []*store.SetupData{
		{Name: userID, Value: 1, IsPrimary: true},
		{Name: userName, Value: "", Size: store.Small},
		{Name: userPass, Value: "", Size: store.Mid},
		{Name: userFirstName, Value: "", Size: store.Small},
		{Name: userLastName, Value: "", Size: store.Small},
		{Name: userEmail, Value: "", Size: store.Mid},
		{Name: userPhone, Value: "", Size: store.Small},
		{Name: userCDate, Value: time.Now()},
		{Name: userExpDate, Value: time.Now().AddDate(2, 0, 0)}}
}

func singleFromStorage(u *User) bool {
	query := store.Get().In(userTable).Attrs(userAttrs...)
	r, _ := query.Equals(userName, u.name).Exec()
	defer r.Close()
	for r.Next() {
		if err := r.Scan(&u.id, &u.name, &u.pass, &u.FirstName, &u.LastName,
			&u.Email, &u.Phone, &u.CDate, &u.ExpDate); err != nil {
			log.DB(err, log.LvlWarn)
			return false
		}
	}
	return true
}

func fromStorageByName(names ...string) []*User {
	query := store.Get().In(userTable).Attrs(userAttrs...)
	if len(names) == 1 {
		query.Equals(userName, names[0])
	} else if len(names) > 1 {
		// query.EqualsMany(userName, names)
	}
	rows, _ := query.Exec()
	return scanUser(rows)
}

func scanUser(rows *sql.Rows) []*User {
	defer rows.Close()
	var us []*User
	for rows.Next() {
		u := User{}
		if err := rows.Scan(&u.id, &u.name, &u.pass, &u.FirstName, &u.LastName,
			&u.Email, &u.Phone, &u.CDate, &u.ExpDate); err != nil {
			log.DB(err, log.LvlWarn)
			return nil
		}
		us = append(us, &u)
	}
	return us
}
