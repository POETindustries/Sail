package pagestore

const pageID = "id"
const pageTitle = "title"
const pageContent = "content"
const pageDomain = "domain"
const pageURL = "url"
const pageStatus = "status"
const pageOwner = "owner"
const pageCreationDate = "cdate"
const pageEditDate = "edate"

const pageAttrs = pageID + "," +
	pageTitle + "," +
	pageContent + "," +
	pageDomain + "," +
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
    ` + pageDomain + ` integer not null default 1,
    ` + pageStatus + ` integer not null default -1,
    ` + pageOwner + ` integer not null default 1,
    ` + pageCreationDate + ` varchar(31) not null default '2015-09-19 10:34:12',
    ` + pageEditDate + ` varchar(31) not null default '2015-09-19 10:34:12');`

const initPage = `do $$ begin if not exists (select id from sl_page)
    then insert into sl_page
    (title, content, url)
    values
    ('Home', 'Welcome to Sail', '/home');
    end if; end $$`
