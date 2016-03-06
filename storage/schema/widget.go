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
const WidgetAttrs = WidgetID + "," + WidgetName + "," + WidgetRefName + "," +
	WidgetType

const createWidget = `create table if not exists sl_widget(
    ` + WidgetID + ` serial not null primary key,
	` + WidgetName + ` varchar(31) not null,
	` + WidgetRefName + ` varchar(31) not null,
    ` + WidgetType + ` varchar(31) not null default 'menu');`

const initWidget = `do $$ begin
	if not exists (select ` + WidgetID + ` from sl_widget)
    then insert into sl_widget
	(` + WidgetName + `,` + WidgetRefName + `)
	values
	('Main Menu', 'main_menu');
    end if;
	end $$`

/********************************************
 * menu widget info
 ********************************************/

const (
	MenuEntryID       = "entry_id"
	MenuWidgetID      = WidgetID
	MenuEntryName     = "entry_name"
	MenuEntryRefID    = "entry_ref_id"
	MenuEntrySubmenu  = "submenu"
	MenuEntryPosition = "position"
)

const MenuAttrs = MenuEntryID + "," + MenuEntryName + "," +
	MenuEntryRefID + "," + MenuEntrySubmenu + "," +
	MenuEntryPosition

const createWidgetMenu = `create table if not exists sl_widget_menu(
    ` + MenuEntryID + ` serial primary key not null,
	` + MenuWidgetID + ` integer not null,
    ` + MenuEntryName + ` varchar(31) not null default 'entry',
    ` + MenuEntryRefID + ` integer not null default 1,
    ` + MenuEntrySubmenu + ` integer not null default 0,
    ` + MenuEntryPosition + ` integer not null default 10);`

const initWidgetMenu = `do $$ begin
	if not exists (select ` + MenuEntryID + ` from sl_widget_menu)
    then insert into sl_widget_menu(
	` + MenuWidgetID + `,
	` + MenuEntryName + `,
	` + MenuEntryRefID + `,
	` + MenuEntryPosition + `)
	values
	(1, 'Home', 1, 0),
	(1, 'About', 4, 10),
	(1, 'Login', 3, 100);
    end if;
	end $$`

/********************************************
 * text widget info
 ********************************************/

const (
	TextID      = WidgetID
	TextContent = "content"
)

const TextAttrs = TextID + "," + TextContent

const createWidgetText = `create table if not exists sl_widget_text(
    ` + TextID + ` integer primary key not null,
    ` + TextContent + ` text not null default '');`
