package page

import (
	"bulma/cachetemplates"
	"errors"
	"io"
	"strings"
	"text/template"
)

type Content struct {
	Title          string
	FontAwesomeURL string
	Body           string
}

type WebPage struct {
	Name         string
	templateName string
	templateHTML []byte

	Content
}

func NewLandingPage(pageName string, templates map[cachetemplates.TemplatePath]cachetemplates.HTML, content Content) (*WebPage, error) {
	page := WebPage{
		Name:         pageName,
		templateName: "page.gohtml",

		Content: content,
	}

	if page.Content.FontAwesomeURL == "" {
		page.Content.FontAwesomeURL = "https://kit.fontawesome.com/908be3e134.js"
	}

	var err error

	page.templateHTML, err = page.getTemplateHTML(templates)
	if err != nil {
		return nil, err
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

func (p *WebPage) getTemplateHTML(templates map[cachetemplates.TemplatePath]cachetemplates.HTML) ([]byte, error) {
	for k, html := range templates {
		if strings.Contains(string(k), p.templateName) {
			return html, nil
		}
	}

	return nil, errors.New("no template was found for page")
}
