package compobody

import (
	"bulma/web"
	"strings"
)

// Body Component
type Body struct {
	templateName string

	markdown []string
}

var _ web.IWeb = (*Body)(nil)

// Template Maybe a constructor should be used.
func (c *Body) Template() string {
	c.templateName = "body.gohtml"

	return c.templateName
}

func (c *Body) AddMarkdown(m string) {
	c.markdown = append(c.markdown, m)
}

func (c *Body) Markdown() string {
	return strings.Join(c.markdown, "")
}
