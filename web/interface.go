package web

import (
	"text/template"
)

// IWeb Groups web components using simple method.
type IWeb interface {
	Render(t *template.Template) (string, error)
}
