package schema

const (
	GroupID              = "group_id"
	GroupName            = "group_name"
	GroupPermMaintenance = "group_perm_maintenance"
	GroupPermUsers       = "group_perm_users"
)

var GroupAttrs = []string{GroupID, GroupName, GroupPermMaintenance,
	GroupPermUsers}

const CreateGroup = `create table if not exists sl_group(
	` + GroupID + ` integer primary key not null,
	` + GroupName + ` text not null,
	` + GroupPermMaintenance + ` integer not null default 0,
	` + GroupPermUsers + ` integer not null default 0);`

const InitGroup = `insert into sl_group values(1,"Admins",7,7);`

const CreateGroupMembers = `create table if not exists sl_group_members(
	` + GroupID + ` integer not null,
	` + UserID + ` integer not null);`

const InitGroupMembers = `insert into sl_group_members values(1,1);`
