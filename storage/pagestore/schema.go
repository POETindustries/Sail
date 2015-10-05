package pagestore

const expDomainID = "domain_id"

const pageID = "page_id"
const pageTitle = "page_title"
const pageContent = "page_content"
const pageDomainID = expDomainID
const pageURL = "page_url"
const pageStatus = "page_status"
const pageOwner = "page_owner"
const pageCreationDate = "page_cdate"
const pageEditDate = "page_edate"

const pageAttrs = pageID + "," +
	pageTitle + "," +
	pageContent + "," +
	pageDomainID + "," +
	pageURL + "," +
	pageStatus + "," +
	pageOwner + "," +
	pageCreationDate + "," +
	pageEditDate

const createPage = `create table if not exists sl_page(
    ` + pageID + ` serial primary key not null,
    ` + pageTitle + ` varchar(63) not null,
    ` + pageURL + ` varchar(63) not null,
    ` + pageContent + ` text not null default '',
    ` + pageDomainID + ` integer not null default 1,
    ` + pageStatus + ` integer not null default -1,
    ` + pageOwner + ` integer not null default 1,
    ` + pageCreationDate + ` varchar(31) not null default '2015-09-19 10:34:12',
    ` + pageEditDate + ` varchar(31) not null default '2015-09-19 10:34:12');`

const initPage = `do $$ begin
	if not exists (select ` + pageID + ` from sl_page)
    then insert into sl_page
    (` + pageTitle + `,` + pageContent + `,` + pageURL + `)
    values
    ('Home', 'Welcome to Sail', '/home');
    end if;
	end $$`

// CreateInstructs contains all database table creation and insert
// instructions that need to be executed when the application starts
// for the first time.
var CreateInstructs = []string{createPage, initPage}
