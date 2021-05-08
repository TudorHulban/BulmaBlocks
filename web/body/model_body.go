package body

import (
	"bulma/cachetemplates"
	"bulma/web"
	"bulma/webcontainers"
	"errors"
	"strings"
	"text/template"
)

// Body Component
type Body struct {
	templateName string
	templateHTML []byte

	Markdown []string
}

var _ web.IWeb = (*Body)(nil)

func NewCo(templates map[cachetemplates.TemplatePath]cachetemplates.HTML) (*Body, error) {
	body := Body{
		templateName: "page.gohtml",
	}

	var err error

	body.templateHTML, err = body.getTemplateHTML(templates)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func (b *Body) getTemplateHTML(templates map[cachetemplates.TemplatePath]cachetemplates.HTML) ([]byte, error) {
	for k, html := range templates {
		if strings.Contains(string(k), b.templateName) {
			return html, nil
		}
	}

	return nil, errors.New("no template was found for page")
}

func (b *Body) AppendToBody(html ...string) {
	b.Markdown = append(b.Markdown, html...)
}

func (b *Body) GetTemplateName() string {
	return b.templateName
}

func (b *Body) Inject(t *template.Template, decorator webcontainers.IWebContainer, blocks ...web.IWeb) error {
	for _, block := range blocks {
		markdown, err := web.Render(t, block)
		if err != nil {
			return err
		}

		b.Markdown = append(b.Markdown, markdown)
	}

	decorator.SetContent(b.Markdown)

	body, err := web.Render(t, decorator)
	if err != nil {
		return err
	}

	b.Markdown = []string{body}

	return nil
}
