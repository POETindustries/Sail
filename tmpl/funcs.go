package tmpl

import (
	"sail/conf"
)

// FetchFiles prepares a slice of template file names for later parsing.
// It expects
func PrepareFiles(files []string) (templates []string) {
	for _, file := range files {
		templates = append(templates, conf.TMPLDIR+file+".html")
	}
	return
}
