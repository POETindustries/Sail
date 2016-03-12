package schema

const (
	PageID         = "page_id"
	PageTitle      = "page_title"
	PageContent    = "page_content"
	PageMetaID     = MetaID
	PageTemplateID = TemplateID
	PageURL        = "page_url"
	PageStatus     = "page_status"
	PageOwner      = "page_owner"
	PageCreateDate = "page_cdate"
	PageEditDate   = "page_edate"
)

const PageAttrs = PageID + "," + PageTitle + "," + PageContent + "," + PageMetaID + "," +
	PageTemplateID + "," + PageURL + "," + PageStatus + "," + PageOwner + "," +
	PageCreateDate + "," + PageEditDate

const CreatePage = `create table if not exists sl_page(
	` + PageID + ` serial primary key not null,
	` + PageTitle + ` varchar(63) not null,
	` + PageURL + ` varchar(63) not null,
	` + PageContent + ` text not null default '',
	` + PageMetaID + ` integer not null default 1,
	` + PageTemplateID + ` integer not null default 1,
	` + PageStatus + ` integer not null default -1,
	` + PageOwner + ` integer not null default 1,
	` + PageCreateDate + ` varchar(31) not null default '2015-09-19 10:34:12',
	` + PageEditDate + ` varchar(31) not null default '2015-09-19 10:34:12');`

const contentLogin = `<form id="login_form" action="/office/" method="POST">
	<input type="text" placeholder="User Name" name="user">
	<input type="password" placeholder="Password" name="pass">
	<input type="submit" value="Submit" id="submit">
</form>`

const InitPage = `do $$ begin
	if not exists (select ` + PageID + ` from sl_page)
    then insert into sl_page(
	` + PageTitle + `,
	` + PageContent + `,
	` + PageURL + `,
	` + PageMetaID + `,
	` + PageTemplateID + `)
	values
	('Home', 'Welcome to Sail', '/home', 1, 1),
	('Office', '', '/office/home', 2, 2),
	('Login', '` + contentLogin + `', '/login', 1, 1),
	('About Sail', 'Go where the wind blows.', '/about', 1, 1);
    end if;
	end $$`
