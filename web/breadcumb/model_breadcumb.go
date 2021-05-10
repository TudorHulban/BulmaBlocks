package breadcumb

import (
	"bulma/cachetemplates"
	"bulma/web"
)

type Content struct {
	ActiveItem string // item presented on the page
	Categories []string
}

// Breadcumb Component
type Breadcumb struct {
	Name         string
	templateName string

	Content
}

var _ web.IRenderer = (*Breadcumb)(nil)

func NewCo(componentName string, templates map[cachetemplates.TemplatePath]cachetemplates.HTML, content Content) *Breadcumb {
	return &Breadcumb{
		Name:         componentName,
		templateName: "breadcrumb.gohtml",

		Content: content,
	}

}

func (b *Breadcumb) TemplateName() string {
	return b.templateName
}
