package schema

const (
	TemplateID       = "template_id"
	TemplateName     = "template_name"
	TemplateWidgetID = WidgetID
)
const TemplateAttrs = TemplateID + "," + TemplateName

const createTemplate = `create table if not exists sl_template(
    ` + TemplateID + ` serial primary key not null,
    ` + TemplateName + ` varchar(31) not null);`

const initTemplate = `do $$ begin
    if not exists(select ` + TemplateID + ` from sl_template)
    then insert into sl_template
    (` + TemplateName + `)
    values
    ('default'),('default-backend');
    end if;
    end $$`

const createTemplateWidgets = `create table if not exists sl_template_widgets(
	` + TemplateID + ` integer not null,
    ` + TemplateWidgetID + ` integer not null);`

const initTemplateWidgets = `do $$ begin
	if not exists(select ` + TemplateID + ` from sl_template_widgets)
	then insert into sl_template_widgets
	(` + TemplateID + `,` + TemplateWidgetID + `) values (1,1);
	end if;
	end $$`
