package domainstore

const expTemplateID = "template_id"

const domainID = "domain_id"
const domainName = "domain_name"
const domainTemplateID = expTemplateID
const domainAttrs = domainID + "," +
	domainName + "," +
	metaAttrs + "," +
	domainTemplateID

const metaTitle = "meta_title"
const metaKeywords = "meta_keywords"
const metaDescription = "meta_description"
const metaLanguage = "meta_language"
const metaPageTopic = "meta_page_topic"
const metaRevisit = "meta_revisit_after"
const metaRobots = "meta_robots"
const metaAttrs = metaTitle + "," +
	metaKeywords + "," +
	metaDescription + "," +
	metaLanguage + "," +
	metaPageTopic + "," +
	metaRevisit + "," +
	metaRobots

const createDomain = `create table if not exists sl_domain(
    ` + domainID + ` serial primary key not null,
    ` + domainName + ` varchar(31) not null,
    ` + metaTitle + ` varchar(63) not null default '',
    ` + metaKeywords + ` varchar(127) not null default'',
    ` + metaDescription + ` text not null default '',
    ` + metaLanguage + ` varchar(31) not null default '',
    ` + metaPageTopic + ` text not null default '',
    ` + metaRevisit + ` varchar(31) not null default '',
    ` + metaRobots + ` varchar(31) not null default '',
    ` + domainTemplateID + ` integer not null default 1);`

const initDomain = `do $$ begin
	if not exists (select ` + domainID + ` from sl_domain)
    then insert into sl_domain values(
    1,
    'default',
    'Sail',
    'cms,content management system,go,golang',
    'Sail is a content management system written in Go',
    'en',
    'cms',
    '1 month',
    'allow',
    1);
    end if;
	end $$`

// CreateInstructs contains all database table creation and insert
// instructions that need to be executed when the application starts
// for the first time.
var CreateInstructs = []string{createDomain, initDomain}
