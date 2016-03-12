package schema

const (
	ContentID         = "content_id"
	ContentTitle      = "content_title"
	ContentContent    = "content_content"
	ContentMetaID     = MetaID
	ContentTemplateID = TemplateID
	ContentURL        = "content_url"
	ContentStatus     = "content_status"
	ContentOwner      = "content_owner"
	ContentCreateDate = "content_cdate"
	ContentEditDate   = "content_edate"
)

var ContentAttrs = []string{ContentID, ContentTitle, ContentContent,
	ContentMetaID, ContentTemplateID, ContentURL, ContentStatus, ContentOwner,
	ContentCreateDate, ContentEditDate}

const CreateContent = `create table if not exists sl_content(
	` + ContentID + ` integer primary key not null,
	` + ContentTitle + ` text not null,
	` + ContentURL + ` text not null,
	` + ContentContent + ` text not null default '',
	` + ContentMetaID + ` integer not null default 1,
	` + ContentTemplateID + ` integer not null default 1,
	` + ContentStatus + ` integer not null default -1,
	` + ContentOwner + ` integer not null default 1,
	` + ContentCreateDate + ` text not null default '2015-09-19 10:34:12',
	` + ContentEditDate + ` text not null default '2015-09-19 10:34:12');`

const contentLogin = `<form id="login_form" action="/office/" method="POST">
	<input type="text" placeholder="User Name" name="user">
	<input type="password" placeholder="Password" name="pass">
	<input type="submit" value="Submit" id="submit">
</form>`

const InitContent = `insert into sl_content(
	` + ContentTitle + `,
	` + ContentContent + `,
	` + ContentURL + `,
	` + ContentMetaID + `,
	` + ContentTemplateID + `)
	values
	('Home', 'Welcome to Sail', '/home', 1, 1),
	('Office', '', '/office/home', 2, 2),
	('Login', '` + contentLogin + `', '/login', 1, 1),
	('About Sail', 'Go where the wind blows.', '/about', 1, 1);`
