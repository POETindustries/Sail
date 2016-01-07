package templatestore

const expWidgetID = "widget_id"

const templateID = "template_id"
const templateName = "template_name"
const templateAttrs = templateID + "," + templateName

const createTemplate = `create table if not exists sl_template(
    ` + templateID + ` serial primary key not null,
    ` + templateName + ` varchar(31) not null);`

const initTemplate = `do $$ begin
    if not exists(select ` + templateID + ` from sl_template)
    then insert into sl_template
    (` + templateName + `)
    values
    ('default'),('default-backend'),('default-login');
    end if;
    end $$`

const createTemplateWidgets = `create table if not exists sl_template_widgets(
	` + templateID + ` integer not null,
    ` + expWidgetID + ` integer not null);`

// CreateInstructs contains all database table creation and insert
// instructions that need to be executed when the application starts
// for the first time.
var CreateInstructs = []string{createTemplate, initTemplate, createTemplateWidgets}
