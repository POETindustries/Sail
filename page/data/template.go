package data

import (
	"fmt"
	"html/template"
	"io"
	"sail/conf"
	"sail/errors"
)

var funcMap = template.FuncMap{
	"even": even}

func even(val int) bool {
	return val%2 == 0
}

// Template is the data structure that contains all data necessary
// to render the template files and all widgets contained within.
type Template struct {
	ID        uint32
	Name      string
	WidgetIDs []uint32
	template  *template.Template
	Widgets   map[string]*Widget
}

// NewTemplate creates a new Template object
func NewTemplate() *Template {
	return &Template{
		Name:    "404",
		Widgets: make(map[string]*Widget)}
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
		tpl, err := template.New("").Funcs(funcMap).ParseGlob(dir + "/*.html")
		if err != nil {
			errors.Log(err, conf.Instance().DevMode)
			tpl, _ = template.New("frame").Parse(NOTFOUND404)
		}
		t.template = tpl
	}
}

// String prints the template's data in an easily readable format.
func (t *Template) String() string {
	str := "TEMPLATE '%s': {ID:%d | WidgetIDs:%+v | Template:%+v | Widgets:%+v}"
	return fmt.Sprintf(str, t.Name, t.ID, t.WidgetIDs, t.template, t.Widgets)
}
