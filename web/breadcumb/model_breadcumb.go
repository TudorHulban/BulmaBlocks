package breadcumb

import (
	"bulma/cachetemplates"
	"bulma/web"
	"errors"
	"io"
	"strings"
	"text/template"
)

type Content struct {
	ActiveItem string // item presented on the page
	Categories []string
}

// Breadcumb Component
type Breadcumb struct {
	Name         string
	templateName string
	templateHTML []byte

	Content
}

var _ web.IWeb = (*Breadcumb)(nil)

func NewCo(breadcumbName string, templates map[cachetemplates.TemplatePath]cachetemplates.HTML, content Content) (*Breadcumb, error) {
	breadcumb := Breadcumb{
		Name:         breadcumbName,
		templateName: "breadcrumb.gohtml",

		Content: content,
	}

	var err error

	breadcumb.templateHTML, err = breadcumb.getTemplateHTML(templates)
	if err != nil {
		return nil, err
	}

	return &breadcumb, nil
}

func (b *Breadcumb) RenderTo(w io.Writer) error {
	t := template.New(b.Name)

	t, errParse := t.Parse(string(b.templateHTML))
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, b)
}

func (b *Breadcumb) TemplateName() string {
	return b.templateName
}

func (b *Breadcumb) templatePath() string {
	return web.TemplateFolderPath + b.templateName
}

func (b *Breadcumb) getTemplateHTML(templates map[cachetemplates.TemplatePath]cachetemplates.HTML) ([]byte, error) {
	for k, html := range templates {
		if strings.Contains(string(k), b.templateName) {
			return html, nil
		}
	}

	return nil, errors.New("no template was found for page")
}
