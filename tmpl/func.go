package tmpl

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"regexp"
	"sail/conf"
	"sail/dbase"
	"strings"
)

func extractNames(tmplContent string) []string {
	re := regexp.MustCompile(`{{template ".*" .*}}`)
	templates := re.FindAllString(tmplContent, -1)

	for i := 0; i < len(templates); i++ {
		templates[i] = strings.TrimPrefix(templates[i], `{{template "`)
		templates[i] = templates[i][:strings.Index(templates[i], `"`)]
	}
	fmt.Println(templates)
	return templates
}

// Builder creates a new Template object and fills it with data as far as that
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
func Builder(db *sql.DB, tmpl string, full bool) *Template {
	t := Template{Files: PrepFiles(subTemplates(db, tmpl))}

	if full {
		t.Name = tmpl
		t.Content = readFile(tmpl)
		t.Templates = subTemplates(db, tmpl)
	}
	return &t
}

// PrepFiles prepares a slice of template names for later parsing.
// It expects a string of template names that are converted to their rooted
// file names on the host os file system. These are ready for parsing by Go's
// own templating engine.
func PrepFiles(files []string) (templates []string) {
	for _, file := range files {
		templates = append(templates, PrepFile(file))
	}
	return
}

// PrepFile prepares one template name  for later parsing.
// It returnes a rooted file name on the host os file system, and expects the
// template directory as stored in the conf package to point to the directory
// that contains all templates on the same level.
func PrepFile(file string) string {
	return conf.New().TmplDir + file + ".html"
}

// ReadFile is a helper function that returns the content of a template file.
// It expects a template name and returns the corresponding file's text or an
// empty string if the file could not be read.
func readFile(tmpl string) string {
	file := PrepFile(tmpl)
	f, err := ioutil.ReadFile(file)

	if err != nil {
		println(err.Error())
		return ""
	}
	return string(f)
}

// subTemplates fetches the template names of all templates used by tmpl.
// It takes a template name, fetches all associated templates from the database
// and returns a slice containing their names. If the template  does not exist
// or anything goes wrong with the database connection, a default 404 template
// is the only template name that the returned slice contains.
func subTemplates(db *sql.DB, tmpl string) []string {
	var csv string
	query := "select templates from sl_template where name=?"
	templates := []string{"404"}

	if row := dbase.QueryRow(query, db, tmpl); row != nil {
		if row.Scan(&csv) == nil {
			templates = strings.Split(csv, ",")
		}
	} // else: templates still contains the 404 template
	return templates
}
