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

var MetaAttrs = []string{MetaTitle, MetaKeywords, MetaDescription,
	MetaLanguage, MetaPageTopic, MetaRevisit, MetaRobots}

const CreateMeta = `create table if not exists sl_meta(
    ` + MetaID + ` integer primary key not null,
    ` + MetaTitle + ` text not null default '',
    ` + MetaKeywords + ` text not null default'',
    ` + MetaDescription + ` text not null default '',
    ` + MetaLanguage + ` text not null default '',
    ` + MetaPageTopic + ` text not null default '',
    ` + MetaRevisit + ` text not null default '',
    ` + MetaRobots + ` text not null default '');`

const InitMeta = `insert into sl_meta values(
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
	'noindex, nofollow');`
