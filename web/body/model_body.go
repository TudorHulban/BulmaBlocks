package body

import (
	"bulma/cachetemplates"
)

// Body Component
type Body struct {
	templateName string
	templateHTML []byte

	Markdown []string
}

func NewCo(templates map[cachetemplates.TemplatePath]cachetemplates.HTML) *Body {
	return &Body{
		templateName: "page.gohtml",
	}
}

// TODO: to remove after adding button at test
func (b *Body) AppendToBody(html ...string) {
	b.Markdown = append(b.Markdown, html...)
}

func (b *Body) Write(markdown []byte) (int, error) {
	b.Markdown = append(b.Markdown, string(markdown))

	return 0, nil
}

func (b *Body) TemplateName() string {
	return b.templateName
}
