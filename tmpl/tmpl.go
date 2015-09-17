package tmpl

import (
	"io/ioutil"
	"regexp"
	"strings"
)

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
func New(name string, full bool) {
	// TODO add actual functionality
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
