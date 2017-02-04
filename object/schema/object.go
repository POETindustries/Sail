package schema

// ObjectID and the other Object[...] constants hold the
// names of the attributes in the object table.
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

// ObjectAttrs is a convenience slice, holding all attribute
// names for the object table. It is intended to be used in
// database queries of the form 'select * from' where the
// order of attributes matters.
var ObjectAttrs = []string{ObjectID, ObjectName, ObjectMachineName,
	ObjectParent, ObjectTypeMajor, ObjectTypeMinor, ObjectStatus, ObjectOwner,
	ObjectCreateDate, ObjectEditDate, ObjectURLCache}

// CreateObject holds the table creation statement.
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

// InitObject executes first-time queries.
const InitObject = `insert into sl_object (
	` + ObjectID + `,
	` + ObjectName + `,
	` + ObjectMachineName + `,
	` + ObjectParent + `,
	` + ObjectTypeMajor + `,
	` + ObjectTypeMinor + `,
	` + ObjectStatus + `,
	` + ObjectURLCache + `
	)
	values
	(1,'Home','home',0,1,0,1,'/home'),
	(2,'About Sail','about',0,1,0,1,'/about'),
	(3,'Files','files',0,0,0,1,'/files/'),
	(4,'Images','img',3,0,0,1,'/files/img/'),
	(5,'Gopher','gopher.png',4,2,1,1,'/files/img/gopher.png');`
