package breadcumb

import (
	"bulma/web"
)

type Content struct {
	ActiveItem string // item presented on the page
	Categories []string
}

// Breadcumb Component
type Breadcumb struct {
	templateName string

	Content
}

var _ web.IWeb = (*Breadcumb)(nil)

func NewCo(item string, categories []string) *Breadcumb {
	return &Breadcumb{
		templateName: "breadcrumb.gohtml",

		Content: Content{
			ActiveItem: item,
			Categories: categories,
		},
	}
}

func (c *Breadcumb) GetTemplateName() string {
	return c.templateName
}

func (c *Breadcumb) getTemplatePath() string {
	return web.TemplateFolderPath + c.templateName
}
