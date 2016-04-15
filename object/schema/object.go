package schema

const (
	ObjectID          = "object_id"
	ObjectRefID       = "object_ref_id"
	ObjectName        = "object_name"
	ObjectMachineName = "object_machine_name"
	ObjectParent      = "object_parent"
	ObjectTypeMajor   = "object_type_major"
	ObjectTypeMinor   = "object_type_minor"
	ObjectStatus      = "object_status"
	ObjectOwner       = "object_owner"
	ObjectCreateDate  = "object_create_date"
	ObjectEditDate    = "object_edit_date"
)

var ObjectAttrs = []string{ObjectID, ObjectRefID, ObjectName, ObjectMachineName,
	ObjectParent, ObjectTypeMajor, ObjectTypeMinor, ObjectStatus, ObjectOwner,
	ObjectCreateDate, ObjectEditDate}

const CreateObject = `create table sl_object(
	` + ObjectID + ` integer primary key not null,
	` + ObjectRefID + ` integer not null default 0,
	` + ObjectName + ` text not null,
	` + ObjectMachineName + ` text not null,
	` + ObjectParent + ` integer not null default 0,
	` + ObjectTypeMajor + ` integer not null default 1,
	` + ObjectTypeMinor + ` integer not null default 1,
	` + ObjectStatus + ` integer not null default 0,
	` + ObjectOwner + ` integer not null default 1,
	` + ObjectCreateDate + ` text not null default '2015-09-19 10:34:12',
	` + ObjectEditDate + ` text not null default '2015-09-19 10:34:12');`

const InitObject = `insert into sl_object (
	` + ObjectRefID + `,
	` + ObjectName + `,
	` + ObjectMachineName + `,
	` + ObjectTypeMajor + `,
	` + ObjectTypeMinor + `,
	` + ObjectStatus + `
	)
	values
	(1,'Home','home',1,1,1),
	(2,'About Sail','about',1,1,1),
	(0,'Gopher','files/img/gopher.png',2,2,1);`
