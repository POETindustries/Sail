package userstore

const userID = "user_id"
const userName = "user_name"
const userPass = "user_pass"
const userFirstName = "user_firstname"
const userLastName = "user_lastname"
const userEmail = "user_email"
const userPhone = "user_phone"
const userCDate = "user_cdate"
const userExpDate = "user_expdate"

const userAttrs = userID + "," +
	userName + "," +
	userPass + "," +
	userFirstName + "," +
	userLastName + "," +
	userEmail + "," +
	userPhone + "," +
	userCDate + "," +
	userExpDate

const createUser = `create table if not exists sl_user(
    ` + userID + ` serial primary key not null,
    ` + userName + ` varchar(31) not null,
    ` + userPass + ` varchar(63) not null,
    ` + userFirstName + ` varchar(63) not null default '',
    ` + userLastName + ` varchar(63) not null default '',
    ` + userEmail + ` varchar(63) not null default '',
    ` + userPhone + ` varchar(31) not null default '',
    ` + userCDate + ` varchar(31) not null default '2015-09-19 10:34:12',
    ` + userExpDate + ` varchar(31) not null default '2020-09-19 10:34:12');`

const initUser = `do $$ begin
    	if not exists (select ` + userID + ` from sl_user)
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

// CreateInstructs contains all database table creation and insert
// instructions that need to be executed when the application starts
// for the first time.
var CreateInstructs = []string{createUser, initUser}
