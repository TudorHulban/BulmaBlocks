package compobody

import (
	"bulma/web"
	"bulma/webcontainers"
	"strings"
	"text/template"
)

// Body Component
type Body struct {
	templateName string

	markdown []string
}

var _ web.IWeb = (*Body)(nil)

func NewCo() *Body {
	return &Body{
		templateName: "body.gohtml",

		markdown: []string{},
	}
}

func (c *Body) GetTemplateName() string {
	return c.templateName
}

func (c *Body) Inject(t *template.Template, decorator webcontainers.IWebContainer, blocks ...web.IWeb) error {
	for _, block := range blocks {
		markdown, err := web.Render(t, block)
		if err != nil {
			return err
		}

		c.markdown = append(c.markdown, markdown)
	}

	decorator.SetContent([]string{c.Markdown()})

	body, err := web.Render(t, decorator)
	if err != nil {
		return err
	}

	//fmt.Println("body:", body)

	c.markdown = []string{body}

	return nil
}

// Markdown Method produces accumulated markdown for body component.
func (c *Body) Markdown() string {
	return strings.Join(c.markdown, "")
}
