package breadcumb

import (
	"bulma/web"
)

type Content struct {
	Item       string
	Categories []string
}

// Breadcumb Component
type Breadcumb struct {
	templateName string

	Content
}

var _ web.IWeb = (*Breadcumb)(nil)

func NewCo(i string, categories []string) *Breadcumb {
	return &Breadcumb{
		templateName: "breadcrumb.gohtml",

		Item:       i,
		Categories: cat,
	}
}

func (c *Breadcumb) GetTemplateName() string {
	return c.templateName
}
