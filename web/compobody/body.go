package compobody

import (
	"strings"
)

// Body Component
type Body struct {
	templateName string

	markdown []string
}

// var _ web.IWeb = (*Body)(nil)

func NewCo() *Body {
	return &Body{
		templateName: "body.gohtml",

		markdown: []string{},
	}
}

func (c *Body) AddMarkdown(m string) {
	c.markdown = append(c.markdown, m)
}

func (c *Body) Markdown() string {
	return strings.Join(c.markdown, "")
}
