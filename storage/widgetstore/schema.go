package widgetstore

const expPageID = "page_id"
const expPageURL = "page_url"

const widgetID = "widget_id"
const widgetName = "widget_name"
const widgetReferenceName = "widget_ref_name"
const widgetType = "widget_type"
const widgetAttrs = widgetID + "," + widgetName + "," +
	widgetReferenceName + "," + widgetType

const menuEntryID = "entry_id"
const menuWidgetID = widgetID
const menuEntryName = "entry_name"
const menuEntryReferenceID = "entry_ref_id"
const menuEntrySubmenu = "submenu"
const menuEntryPosition = "position"
const menuAttrs = menuEntryID + "," + menuEntryName + "," +
	menuEntryReferenceID + "," + menuEntrySubmenu + "," +
	menuEntryPosition

const textID = widgetID
const textContent = "content"
const textAttrs = textID + "," + textContent

const createWidget = `create table if not exists sl_widget(
    ` + widgetID + ` serial not null primary key,
	` + widgetName + ` varchar(31) not null,
	` + widgetReferenceName + ` varchar(31) not null,
    ` + widgetType + ` varchar(31) not null default 'menu');`

const createWidgetMenu = `create table if not exists sl_widget_menu(
    ` + menuEntryID + ` serial primary key not null,
	` + menuWidgetID + ` integer not null,
    ` + menuEntryName + ` varchar(31) not null default 'entry',
    ` + menuEntryReferenceID + ` integer not null default 1,
    ` + menuEntrySubmenu + ` integer not null default 0,
    ` + menuEntryPosition + ` integer not null default 10);`

const createWidgetText = `create table if not exists sl_widget_text(
    ` + textID + ` integer primary key not null,
    ` + textContent + ` text not null default '');`

// CreateInstructs contains all database table creation and insert
// instructions that need to be executed when the application starts
// for the first time.
var CreateInstructs = []string{createWidget, createWidgetMenu, createWidgetText}
