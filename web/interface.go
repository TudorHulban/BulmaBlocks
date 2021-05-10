package web

import (
	"bulma/cachetemplates"
	"errors"
	"io"
	"strings"
	"text/template"
)

// IRenderer Groups web components using simple method.
type IRenderer interface {
	TemplateName() string
}

// type URL string

const TemplateFolderPath = "../../templates/"

func templateHTML(webComponent IRenderer, templates map[cachetemplates.TemplatePath]cachetemplates.HTML) ([]byte, error) {
	for k, html := range templates {
		if strings.Contains(string(k), webComponent.TemplateName()) {
			return html, nil
		}
	}

	return nil, errors.New("no template was found for page")
}

func render(w io.Writer, templateHTML []byte, webComponent IRenderer) error {
	t := template.New("")

	t, errParse := t.Parse(string(templateHTML))
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, webComponent)
}

func RenderComponentTo(webComponent IRenderer, to io.Writer, templates map[cachetemplates.TemplatePath]cachetemplates.HTML) error {
	html, errCache := templateHTML(webComponent, templates)
	if errCache != nil {
		return errCache
	}

	return render(to, html, webComponent)
}
