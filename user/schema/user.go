package schema

const (
	UserID        = "user_id"
	UserName      = "user_name"
	UserPass      = "user_pass"
	UserFirstName = "user_firstname"
	UserLastName  = "user_lastname"
	UserEmail     = "user_email"
	UserPhone     = "user_phone"
	UserCDate     = "user_cdate"
	UserExpDate   = "user_expdate"
)

var UserAttrs = []string{UserID, UserName, UserPass, UserFirstName,
	UserLastName, UserEmail, UserPhone, UserCDate, UserExpDate}

const CreateUser = `create table if not exists sl_user(
    ` + UserID + ` integer primary key not null,
    ` + UserName + ` text not null,
    ` + UserPass + ` text not null,
    ` + UserFirstName + ` text not null default '',
    ` + UserLastName + ` text not null default '',
    ` + UserEmail + ` text not null default '',
    ` + UserPhone + ` text not null default '',
    ` + UserCDate + ` text not null default '2015-09-19 10:34:12',
    ` + UserExpDate + ` text not null default '2020-09-19 10:34:12');`

// InitUser inserts a default user with admin rights. Password is 'toor'.
// Obviously not suitable for production environments.
const InitUser = `insert into sl_user
        values(
        1,
        'admin',
		'$2a$12$gZaLvx7eUgn7DA0XDwKgOe//FvyXO7RynueOeXOnQrCkmwr/Y0rRS',
        'Mister',
        'Goodkat',
        'admin@sail.example.com',
        '',
        '2015-09-19 10:34:12',
        '2020-09-19 10:34:12');`
