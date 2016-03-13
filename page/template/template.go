package template

import (
	"fmt"
	"html/template"
	"io"
	"sail/conf"
	"sail/errors"
	"sail/page/fallback"
	"sail/page/widget"
)

// Template is the data structure that contains all data necessary
// to render the template files and all widgets contained within.
type Template struct {
	ID        uint32
	Name      string
	WidgetIDs []uint32
	template  *template.Template
	Widgets   map[string]*widget.Widget
}

// New creates a new Template object
func New() *Template {
	return &Template{
		Name:    "404",
		Widgets: make(map[string]*widget.Widget)}
}

// ByID returns the template that matches the given id.
//
// It should be used to prepare one template for rendering
// and is guaranteed to return a pointer to a valid template.
func ByID(id uint32) *Template {
	ts := fromStorageByID(id)
	if len(ts) < 1 {
		if id == 1 {
			return New()
		}
		return ByID(1)
	}
	return ts[0]
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

// Parse parses the template files pointed at by the template.
func (t *Template) Parse() {
	if t.Name == "404" {
		t.template, _ = template.New("frame").Parse(fallback.NOTFOUND404)
	} else {
		dir := conf.Instance().TmplDir + t.Name
		tpl, err := template.New("").Funcs(funcMap).ParseGlob(dir + "/*.html")
		if err != nil {
			errors.Log(err, conf.Instance().DevMode)
			tpl, _ = template.New("frame").Parse(fallback.NOTFOUND404)
		}
		t.template = tpl
	}
}

// String prints the template's data in an easily readable format.
func (t *Template) String() string {
	str := "TEMPLATE '%s': {ID:%d | WidgetIDs:%+v | Template:%+v | Widgets:%+v}"
	return fmt.Sprintf(str, t.Name, t.ID, t.WidgetIDs, t.template, t.Widgets)
}
