package tmpl

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sail/conf"
	"strings"
)

var files = [1]string{"menu"}

// FetchFiles prepares a slice of template file names for later parsing.
func FetchFiles(frame string) *[]string {
	var templates []string
	templates = append(templates, conf.TMPLDIR+frame+".html")
	for _, file := range files {
		templates = append(templates, conf.TMPLDIR+file+".html")
	}
	readFile(templates[0])
	return &templates
}

func readFile(file string) {
	f, _ := ioutil.ReadFile(file) // TODO catch error return value
	re := regexp.MustCompile(`{{template ".*" .*}}`)
	templates := re.FindAllString(string(f), -1)

	for i := 0; i < len(templates); i++ {
		templates[i] = strings.TrimPrefix(templates[i], "{{template \"")
		templates[i] = templates[i][:strings.Index(templates[i], "\"")]
	}

	fmt.Println(templates)
}
