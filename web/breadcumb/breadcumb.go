package breadcumb

import (
	"bulma/web"
	"text/template"
)

// Breadcumb Component
type Breadcumb struct {
	templateName string

	Item       string
	Categories []string
}

var _ web.IWeb = (*Breadcumb)(nil)

func NewCo(i string, cat []string) *Breadcumb {
	return &Breadcumb{
		templateName: "breadcrumb.gohtml",

		Item:       i,
		Categories: cat,
	}
}

func (c *Breadcumb) Render(t *template.Template) (string, error) {
	return web.Render(t, c.templateName, c)
}
