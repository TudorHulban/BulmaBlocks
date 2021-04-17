package compobody

import (
	"bulma/web"
	"strings"
	"text/template"
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

func (c *Body) Inject(t *template.Template, blocks ...web.IWeb) error {
	for _, block := range blocks {
		markdown, err := block.Render(t)
		if err != nil {
			return err
		}

		c.markdown = append(c.markdown, markdown)
	}

	return nil
}

func (c *Body) Markdown() string {
	return strings.Join(c.markdown, "")
}
