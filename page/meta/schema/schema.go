package schema

const (
	MetaID          = "meta_id"
	MetaTitle       = "meta_title"
	MetaKeywords    = "meta_keywords"
	MetaDescription = "meta_description"
	MetaLanguage    = "meta_language"
	MetaPageTopic   = "meta_page_topic"
	MetaRevisit     = "meta_revisit_after"
	MetaRobots      = "meta_robots"
)

const MetaAttrs = MetaTitle + "," + MetaKeywords + "," + MetaDescription + "," +
	MetaLanguage + "," + MetaPageTopic + "," + MetaRevisit + "," + MetaRobots

const CreateMeta = `create table if not exists sl_meta(
    ` + MetaID + ` serial primary key not null,
    ` + MetaTitle + ` varchar(63) not null default '',
    ` + MetaKeywords + ` varchar(127) not null default'',
    ` + MetaDescription + ` text not null default '',
    ` + MetaLanguage + ` varchar(31) not null default '',
    ` + MetaPageTopic + ` text not null default '',
    ` + MetaRevisit + ` varchar(31) not null default '',
    ` + MetaRobots + ` varchar(31) not null default '');`

const InitMeta = `do $$ begin
	if not exists (select ` + MetaID + ` from sl_meta)
    then insert into sl_meta values(
    1,
    'Sail',
    'cms,content management system,go,golang',
    'Sail is a content management system written in Go',
    'en',
    'cms',
    '1 month',
    'index, follow'),(
	2,
	'Sail Backend',
	'',
	'',
	'',
	'',
	'',
	'noindex, nofollow');
    end if;
	end $$`
