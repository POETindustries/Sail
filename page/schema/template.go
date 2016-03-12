package schema

const (
	TemplateID       = "template_id"
	TemplateName     = "template_name"
	TemplateWidgetID = WidgetID
)

var TemplateAttrs = [...]string{TemplateID, TemplateName}

const CreateTemplate = `create table if not exists sl_template(
    ` + TemplateID + ` integer primary key not null,
    ` + TemplateName + ` text not null);`

const InitTemplate = `insert into sl_template
    (` + TemplateName + `)
    values
    ('default'),('default-backend');`

const CreateTemplateWidgets = `create table if not exists sl_template_widgets(
	` + TemplateID + ` integer not null,
    ` + TemplateWidgetID + ` integer not null);`

const InitTemplateWidgets = `insert into sl_template_widgets
	(` + TemplateID + `,` + TemplateWidgetID + `) values (1,1);`
