package compobody

import (
	"bulma/web"
	"bulma/web/container"
	"strings"
	"text/template"
)

// Body Component
type Body struct {
	templateName string

	markdown []string
}

func NewCo() *Body {
	return &Body{
		templateName: "body.gohtml",

		markdown: []string{},
	}
}

func (c *Body) Inject(t *template.Template, decorator *container.Container, blocks ...web.IWeb) error {
	for _, block := range blocks {
		markdown, err := block.Render(t)
		if err != nil {
			return err
		}

		c.markdown = append(c.markdown, markdown)
	}

	decorator.SetMarkdonw(c.Markdown())

	body, err := decorator.Render(t)
	if err != nil {
		return err
	}

	c.markdown = []string{body}

	return nil
}

// Markdown Method produces accumulated markdown for body component.
func (c *Body) Markdown() string {
	return strings.Join(c.markdown, "")
}
