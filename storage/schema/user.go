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

const UserAttrs = UserID + "," + UserName + "," + UserPass + "," +
	UserFirstName + "," + UserLastName + "," + UserEmail + "," +
	UserPhone + "," + UserCDate + "," + UserExpDate

const createUser = `create table if not exists sl_user(
    ` + UserID + ` serial primary key not null,
    ` + UserName + ` varchar(31) not null,
    ` + UserPass + ` varchar(63) not null,
    ` + UserFirstName + ` varchar(63) not null default '',
    ` + UserLastName + ` varchar(63) not null default '',
    ` + UserEmail + ` varchar(63) not null default '',
    ` + UserPhone + ` varchar(31) not null default '',
    ` + UserCDate + ` varchar(31) not null default '2015-09-19 10:34:12',
    ` + UserExpDate + ` varchar(31) not null default '2020-09-19 10:34:12');`

const initUser = `do $$ begin
    	if not exists (select ` + UserID + ` from sl_user)
        then insert into sl_user
        values(
        1,
        'admin',
		'$2a$08$g312eq5uYaPXoixBVkOI4OC2Su1sgS2GOhOuNPJqMBrn1j/99vKe2',
        '',
        '',
        'admin@sail.example.com',
        '',
        '2015-09-19 10:34:12',
        '2020-09-19 10:34:12');
        end if;
    	end $$`
