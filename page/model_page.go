package page

import (
	"bulma/cachetemplates"
	"errors"
	"io"
	"strings"
	"text/template"
)

type Content struct {
	Author string
}

type WebPage struct {
	Name         string
	templateName string
	templateHTML []byte

	Content
}

func NewLandingPage(pageName string, content Content, templates map[cachetemplates.TemplatePath]cachetemplates.HTML) (*WebPage, error) {
	page := WebPage{
		Name:         pageName,
		templateName: "page.gohtml",

		Content: content,
	}

	return &page, nil
}

func (p *WebPage) RenderTo(w io.Writer) error {
	t := template.New(p.Name)

	t, errParse := t.Parse(string(p.templateHTML))
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, p)
}

func (p *WebPage) getTemplateHTML(templates map[cachetemplates.TemplatePath]cachetemplates.HTML) (string, error) {
	for k, v := range templates {
		if strings.Contains(string(k), p.templateName) {
			return string(v), nil
		}
	}

	return "", errors.New("no template was found for page")
}
