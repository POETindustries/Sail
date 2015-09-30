package storage

const PageID = "id"
const pageTitle = "title"
const pageContent = "content"
const pageDomain = "domain"
const PageURL = "url"
const pageStatus = "status"
const pageOwner = "owner"
const pageCreationDate = "cdate"
const pageEditDate = "edate"

const PageAttrs = PageID + "," +
	pageTitle + "," +
	pageContent + "," +
	pageDomain + "," +
	PageURL + "," +
	pageStatus + "," +
	pageOwner + "," +
	pageCreationDate + "," +
	pageEditDate

const domainID = "id"
const domainName = "name"
const domainTemplateID = "template_id"
const DomainAttrs = domainID + "," +
	domainName + "," +
	MetaAttrs + "," +
	domainTemplateID

const metaTitle = "title"
const metaKeywords = "keywords"
const metaDescription = "description"
const metaLanguage = "language"
const metaPageTopic = "page_topic"
const metaRevisit = "revisit_after"
const metaRobots = "robots"
const MetaAttrs = metaTitle + "," +
	metaKeywords + "," +
	metaDescription + "," +
	metaLanguage + "," +
	metaPageTopic + "," +
	metaRevisit + "," +
	metaRobots

const templateID = "id"
const templateName = "name"
const TemplateAttrs = templateID + "," + templateName

const widgetID = "id"
const widgetName = "name"
const widgetType = "type"
const WidgetAttrs = widgetID + "," + widgetName + "," + widgetType

const textWidgetID = widgetID
const textWidgetContent = "content"
const TextWidgetAttrs = textWidgetID + "," + textWidgetContent

const menuWidgetEntryID = "id"
const menuWidgetEntryName = "entry_name"
const menuWidgetEntryReferenceID = "ref_id"
const menuWidgetEntrySubmenu = "has_submenu"
const menuWidgetEntryPosition = "position"
const MenuWidgetAttrs = menuWidgetEntryID + "," +
	menuWidgetEntryName + "," +
	menuWidgetEntryReferenceID + "," +
	menuWidgetEntrySubmenu + "," +
	menuWidgetEntryPosition

const createPage = `create table if not exists sl_page(
    ` + PageID + ` serial primary key not null,
    ` + pageTitle + ` varchar(63) not null,
    ` + PageURL + ` varchar(63) not null,
    ` + pageContent + ` text not null default '',
    ` + pageDomain + ` integer not null default 1,
    ` + pageStatus + ` integer not null default -1,
    ` + pageOwner + ` integer not null default 1,
    ` + pageCreationDate + ` varchar(31) not null default '2015-09-19 10:34:12',
    ` + pageEditDate + ` varchar(31) not null default '2015-09-19 10:34:12');`

const createDomain = `create table if not exists sl_domain(
    ` + domainID + ` serial primary key not null,
    ` + domainName + ` varchar(31) not null,
    ` + metaTitle + ` varchar(63) not null default '',
    ` + metaKeywords + ` varchar(127) not null default'',
    ` + metaDescription + ` text not null default '',
    ` + metaLanguage + ` varchar(31) not null default '',
    ` + metaPageTopic + ` text not null default '',
    ` + metaRevisit + ` varchar(31) not null default '',
    ` + metaRobots + ` varchar(31) not null default ''
    ` + domainTemplateID + ` integer not null default 1);`

const createTemplate = `create table if not exists sl_template(
    ` + templateID + ` serial primary key not null,
    ` + templateName + ` varchar(31) not null);`

const createWidget = `create table if not exists sl_widget(
    ` + widgetID + ` serial not null primary key,
    ` + widgetName + ` varchar(31) not null,
    ` + widgetType + ` varchar(31) not null default 'menu');`

const createWidgetTextField = `create table if not exists sl_widget_text(
    ` + textWidgetID + ` integer primary key not null,
    ` + textWidgetContent + ` text not null default '');`

const initPage = `do $$ begin if not exists (select id from sl_page)
    then insert into sl_page
    (title, content, url)
    values
    ('Home', 'Welcome to Sail', '/home');
    end if; end $$`

const initDomain = `do $$ begin if not exists (select id from sl_domain)
    then insert into sl_domain values(
    1,
    'default',
    'Sail',
    'cms,content management system, go, golang',
    'Sail is a content management system written in Go',
    'en',
    'cms',
    '1 month',
    'allow',
    1);
    end if; end $$`

const initTemplate = `do $$ begin if not exists(select id from sl_template)
    then insert into sl_template
    (name)
    values
    ('default');
    end if; end $$`

var createInstructs = []string{
	createWidget,
	createWidgetTextField,
	createPage,
	createDomain,
	createTemplate,
	createTemplateTable(1),
	initPage,
	initDomain,
	initTemplate}

func createTemplateTable(id int) string {
	return `create table if not exists sl_template_` + string(id) + `(
        ` + widgetID + ` integer primary key not null);`
}

func deleteTemplateTable(id int) string {
	return `drop table if exists sl_template_` + string(id) + ";"
}

func createMenuWidgetTable(id int) string {
	return `create table if not exists sl_widget_menu_` + string(id) + `(
        ` + menuWidgetEntryID + ` serial primary key not null,
        ` + menuWidgetEntryName + ` varchar(31) not null default 'entry',
        ` + menuWidgetEntryReferenceID + ` integer not null default 1,
        ` + menuWidgetEntrySubmenu + ` boolean not null default false,
        ` + menuWidgetEntryPosition + ` integer not null default 100);`
}

func deleteMenuWidgetTable(id int) string {
	return `drop table if exists sl_widget_menu_` + string(id) + ";"
}
