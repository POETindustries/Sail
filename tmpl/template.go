package tmpl

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type Template struct {
	Templates []string
	Files     []string
	Content   string
	FileName  string
}

func readFile(file string) string {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		println(err.Error())
		return ""
	}
	return string(f)
}

func extractNames(tmpl string) []string {
	re := regexp.MustCompile(`{{template ".*" .*}}`)
	templates := re.FindAllString(tmpl, -1)

	for i := 0; i < len(templates); i++ {
		templates[i] = strings.TrimPrefix(templates[i], `{{template "`)
		templates[i] = templates[i][:strings.Index(templates[i], `"`)]
	}

	fmt.Println(templates)
	return templates
}
