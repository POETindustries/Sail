package template

import "html/template"

var funcMap = template.FuncMap{
	"even": even}

func even(val int) bool {
	return val%2 == 0
}
