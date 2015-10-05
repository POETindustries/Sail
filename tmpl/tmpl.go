package tmpl

import (
	"fmt"
	"html/template"
	"io"
	"sail/conf"
	"sail/errors"
	"sail/widget"
)

// NOTFOUND404 is a very basic web page signaling a 404 error.
// It contails the bare minimum necessary for a syntactically correct html web
// page and is used in those cases when not even basic database connections
// and templates work. The cms cannot be considered functional should that
// happen, and this markup at least tells the user as much. The markup is as
// generic as possible while still being somewhat good looking.
const NOTFOUND404 = `
	<head><title>Sorry About That</title><meta charset="utf-8"></head>
	<body style="background:black;text-align:center;color:white;padding:72px;font-size:1.5em;">
		<p style="font-size:2em;">Sorry About That!</p>
		<p>PAGE NOT FOUND</p>
	</body>
`

// Template is the data structure that contains all data necessary
// to render the template files and all widgets contained within.
type Template struct {
	ID        uint32
	Name      string
	WidgetIDs []uint32
	template  *template.Template
	widgets   map[string]*widget.Widget
}

// Execute applies a parsed template to the specified data object,
// writing the output to wr. If an error occurs during execution, it
// is the responsibility of the caller to handle partially written
// output.
func (t *Template) Execute(wr io.Writer, data interface{}) (err error) {
	if t.template == nil {
		err = errors.NilPointer()
	} else {
		err = t.template.ExecuteTemplate(wr, "frame.html", data)
	}
	errors.Log(err, conf.Instance().DevMode)
	return
}

// Compile parses the template files pointed at by the template.
func (t *Template) Compile() {
	if t.Name == "404" {
		t.template, _ = template.New("frame").Parse(NOTFOUND404)
	} else {
		dir := conf.Instance().TmplDir + t.Name
		tpl, err := template.ParseGlob(dir + "/*.html")
		if err != nil {
			errors.Log(err, conf.Instance().DevMode)
			tpl, _ = template.New("frame").Parse(NOTFOUND404)
		}
		t.template = tpl
	}
}

// Widget returns a pointer to the widget designated by the name
// parameter. If no such widget exists,
func (t *Template) Widget(name string) (w *widget.Widget) {
	if w = t.widgets[name]; w == nil {
		return widget.New()
	}
	return
}

// AddWidget adds the widget w to the template's internal collection
// of widgets. If the widget or a widget with the same name already
// exists, it is overwritten.
func (t *Template) AddWidget(w *widget.Widget) {
	t.widgets[w.RefName] = w
}

// String prints the template's data in an easily readable format.
func (t *Template) String() string {
	str := "TEMPLATE '%s': {ID:%d | WidgetIDs:%+v | Template:%+v | Widgets:%+v}"
	return fmt.Sprintf(str, t.Name, t.ID, t.WidgetIDs, t.template, t.widgets)
}

// New creates a new Template object
func New() *Template {
	return &Template{
		Name:    "404",
		widgets: make(map[string]*widget.Widget)}
}
