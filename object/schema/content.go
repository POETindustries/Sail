package schema

// ContentID, ContentContent, ContentMetaID and ContentTemplateID
// hold the names for the table's columns.
const (
	ContentID         = ObjectID
	ContentContent    = "content_content"
	ContentMetaID     = MetaID
	ContentTemplateID = TemplateID
)

// ContentAttrs holds all the table's attributes for convenient
// use in queries where all columns need to be fetched.
var ContentAttrs = []string{ContentID, ContentContent, ContentMetaID,
	ContentTemplateID}

// CreateContent holds the table creation statement as it
// would be called upon database creation.
const CreateContent = `create table if not exists sl_content(
	` + ContentID + ` integer primary key not null,
	` + ContentContent + ` text not null default '',
	` + ContentMetaID + ` integer not null default 1,
	` + ContentTemplateID + ` integer not null default 1);`

// InitContent fills the empty table with the minimum amount
// of data that needs to be present in order for Sail to function.
const InitContent = `insert into sl_content(
	` + ContentID + `,
	` + ContentContent + `,
	` + ContentMetaID + `,
	` + ContentTemplateID + `)
	values
	(1, '<p>Welcome to Sail</p><img width="200px" src="uuid/5"/>', 1, 1),
	(2, '<p>Go where the wind blows.</p>', 1, 1);`
