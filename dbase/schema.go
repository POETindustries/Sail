package dbase

const createPage = `create table if not exists sl_page(
    id serial primary key not null,
    title varchar(63) not null,
    content text not null default '',
    domain integer not null default 1,
    url varchar(63) not null);`

const createPageMeta = `create table if not exists sl_page_meta(
    id integer not null unique,
    title varchar(63) not null default '',
    keywords varchar(127) not null default'',
    description text not null default '',
    language varchar(31) not null default '',
    page_topic text not null default '',
    revisit_after varchar(31) not null default '',
    robots varchar(31) not null default '');`

const createDomain = `create table if not exists sl_domain(
    id serial primary key not null,
    name varchar(31) not null,
    template text not null default 'default');`

const initPage = `insert into sl_page
    (title, content, url)
    values
    ('Home', 'Welcome to Sail', '/home')`

const initDomain = `insert into sl_domain
    (name, template)
    values
    ('default', 'default');`

const initMeta = `insert into sl_page_meta values(
    1,
    'Sail',
    'cms,content management system, go, golang',
    'Sail is a content management system written in Go',
    'en',
    'cms',
    '1 month',
    'allow');`

var createInstructs = [...]string{
	createPage,
	createPageMeta,
	createDomain,
	initPage,
	initDomain,
	initMeta}
