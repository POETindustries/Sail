package tmpl

import (
	"html/template"
	"io"
	"io/ioutil"
	"regexp"
	"sail/conf"
	"sail/errors"
	"strings"
)

// NOTFOUND404 is a very basic web page signaling a 404 error.
// It contails the bare minimum necessary for a syntactically correct html web
// page and is used in those cases when not even basic database connections
// and templates work. The cms cannot be considered functional should that
// happen, and this markup at least tells the user as much. The markup is as
// generic as possible while still being somewhat good looking.
const NOTFOUND404 = `<!doctype html>
		<html style="background:black;text-align:center;color:white;">
		<head><title>Sorry About That</title><meta charset="utf-8"></head>
		<body style="padding:72px;font-family:sans-serif;font-size:1.5em;">
		<p style="font-size:2em;">Sorry About That!</p>
		<p>PAGE NOT FOUND</p></body></html>`

type Template struct {
	template *template.Template
	widgets  map[string]interface{}
}

func (t *Template) Execute(wr io.Writer, data interface{}) error {
	err := t.template.ExecuteTemplate(wr, "frame.html", data)

	if err != nil {
		errors.Log(err, conf.Instance().DevMode)
	}

	return err
}

func (t *Template) Load(name string) {
	if name == "404" {
		t.template, _ = template.New("frame").Parse(NOTFOUND404)
	} else {
		dir := conf.Instance().TmplDir + name

		tpl, err := template.ParseGlob(dir + "/*.html")
		if err != nil {
			errors.Log(err, conf.Instance().DevMode)
			tpl, _ = template.New("frame").Parse(NOTFOUND404)
		}

		t.template = tpl
	}
}

func (t *Template) Widget(name string) interface{} {
	return t.widgets[name]
}

// New creates a new Template object and fills it with data as far as that
// data exists.
//
// The 'full' flag allows us to specify if we want all data to be loaded or if
// zero values suffice for the current use case. The reasoning behind this is
// that for frontend page building only the template file names are necessary.
// Thus, most of the time we only need the Files field to contain meaningful
// and correct data.
//
// The values of the other fields are only used when editing templates, which
// happens orders of magnitude less frequent than simple display for the
// average page visitor. Only then is there a need for a completely populated
// and large struct.
func New(name string) *Template {
	t := Template{}
	t.Load(name)
	/*	t.widgets = make(map[string]interface{})
		for _, id := range widgetIDs {
			widget := widget.Menu{}
			widget.ScanFromDB("id", id)
			t.widgets[widget.Name] = widget
		}*/
	return &t
}

// ReadFile is a helper function that returns the content of a template file.
// It expects a template name and returns the corresponding file's text or an
// empty string if the file could not be read.
func ReadFile(file string) string {
	if f, err := ioutil.ReadFile(file); err == nil {
		return string(f)
	}

	return ""
}

func subTemplates(tmplContent string) []string {
	re := regexp.MustCompile(`{{template ".*" .*}}`)
	templates := re.FindAllString(tmplContent, -1)

	for i := 0; i < len(templates); i++ {
		templates[i] = strings.TrimPrefix(templates[i], `{{template "`)
		templates[i] = templates[i][:strings.Index(templates[i], `"`)]
	}

	return templates
}
