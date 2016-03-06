package schema

// CreateInstructs contains all database table creation and insert
// instructions that need to be executed when the application starts
// for the first time.
var CreateInstructs = []string{
	createUser,
	initUser,
	createWidget,
	initWidget,
	createWidgetMenu,
	initWidgetMenu,
	createWidgetText,
	createTemplate,
	initTemplate,
	createTemplateWidgets,
	initTemplateWidgets,
	createMeta,
	initMeta,
	createPage,
	initPage}
