package schema

/********************************************
 * generic widget info
 ********************************************/

// WidgetID holds the unique widget id
const (
	WidgetID      = "widget_id"
	WidgetName    = "widget_name"
	WidgetRefName = "widget_ref_name"
	WidgetType    = "widget_type"
)

// WidgetAttrs is a convenient string for queries of the type
// 'select * from'.
var WidgetAttrs = []string{WidgetID, WidgetName, WidgetRefName, WidgetType}

const CreateWidget = `create table if not exists sl_widget(
    ` + WidgetID + ` integer not null primary key,
	` + WidgetName + ` text not null,
	` + WidgetRefName + ` text not null,
    ` + WidgetType + ` text not null default 'nav');`

const InitWidget = `insert into sl_widget
	(` + WidgetName + `,` + WidgetRefName + `)
	values
	('Main Menu', 'main_menu');`

/********************************************
 * nav widget info
 ********************************************/

const (
	NavEntryID       = "entry_id"
	NavWidgetID      = WidgetID
	NavEntryName     = "entry_name"
	NavEntryRefID    = ContentID
	NavEntrySubmenu  = "submenu"
	NavEntryPosition = "position"
)

var NavAttrs = []string{NavEntryID, NavEntryName, NavEntryRefID,
	NavEntrySubmenu, NavEntryPosition}

const CreateWidgetNav = `create table if not exists sl_widget_nav(
    ` + NavEntryID + ` integer primary key not null,
	` + NavWidgetID + ` integer not null,
    ` + NavEntryName + ` text not null default 'entry',
    ` + NavEntryRefID + ` integer not null default 1,
    ` + NavEntrySubmenu + ` integer not null default 0,
    ` + NavEntryPosition + ` integer not null default 10);`

const InitWidgetNav = `insert into sl_widget_nav(
	` + NavWidgetID + `,
	` + NavEntryName + `,
	` + NavEntryRefID + `,
	` + NavEntryPosition + `)
	values
	(1, 'Home', 1, 0),
	(1, 'About', 2, 10);`

/********************************************
 * text widget info
 ********************************************/

const (
	TextID      = WidgetID
	TextContent = "content"
)

var TextAttrs = []string{TextID, TextContent}

const CreateWidgetText = `create table if not exists sl_widget_text(
    ` + TextID + ` integer not null,
    ` + TextContent + ` text not null default '');`
