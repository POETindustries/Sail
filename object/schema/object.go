package schema

const (
	ObjectID          = "object_id"
	ObjectName        = "object_name"
	ObjectMachineName = "object_machine_name"
	ObjectParent      = "object_parent"
	ObjectTypeMajor   = "object_type_major"
	ObjectTypeMinor   = "object_type_minor"
	ObjectStatus      = "object_status"
	ObjectOwner       = "object_owner"
	ObjectCreateDate  = "object_create_date"
	ObjectEditDate    = "object_edit_date"
	ObjectURLCache    = "object_url_cache"
)

var ObjectAttrs = []string{ObjectID, ObjectName, ObjectMachineName,
	ObjectParent, ObjectTypeMajor, ObjectTypeMinor, ObjectStatus, ObjectOwner,
	ObjectCreateDate, ObjectEditDate, ObjectURLCache}

const CreateObject = `create table sl_object(
	` + ObjectID + ` integer primary key not null,
	` + ObjectName + ` text not null,
	` + ObjectMachineName + ` text not null,
	` + ObjectParent + ` integer not null default 0,
	` + ObjectTypeMajor + ` integer not null default 1,
	` + ObjectTypeMinor + ` integer not null default 1,
	` + ObjectStatus + ` integer not null default 0,
	` + ObjectOwner + ` integer not null default 1,
	` + ObjectCreateDate + ` text not null default '2015-09-19 10:34:12',
	` + ObjectEditDate + ` text not null default '2015-09-19 10:34:12',
	` + ObjectURLCache + ` text not null default '');`

const InitObject = `insert into sl_object (
	` + ObjectName + `,
	` + ObjectMachineName + `,
	` + ObjectTypeMajor + `,
	` + ObjectTypeMinor + `,
	` + ObjectStatus + `,
	` + ObjectURLCache + `
	)
	values
	('Home','home',1,1,1, '/home'),
	('About Sail','about',1,1,1, '/about'),
	('Gopher','files/img/gopher.png',2,2,1, '/files/img/gopher.png');`
