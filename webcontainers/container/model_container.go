package container

import (
	"bulma/cachetemplates"
	"errors"
	"io"
	"strings"
	"text/template"
)

// Container Component
type Container struct {
	Name         string
	templateName string
	templateHTML []byte

	Markdown []string
}

func NewCo(containerName string, templates map[cachetemplates.TemplatePath]cachetemplates.HTML) (*Container, error) {
	container := Container{
		Name:         containerName,
		templateName: "container.gohtml",
	}

	var err error

	container.templateHTML, err = container.getTemplateHTML(templates)
	if err != nil {
		return nil, err
	}

	return &container, nil
}

func (c *Container) Write(markdown []byte) (int, error) {
	c.Markdown = append(c.Markdown, string(markdown))

	return 0, nil
}

func (c *Container) RenderTo(w io.Writer) error {
	t := template.New(c.Name)

	t, errParse := t.Parse(string(c.templateHTML))
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, c)
}

func (c *Container) getTemplateHTML(templates map[cachetemplates.TemplatePath]cachetemplates.HTML) ([]byte, error) {
	for k, html := range templates {
		if strings.Contains(string(k), c.templateName) {
			return html, nil
		}
	}

	return nil, errors.New("no template was found for page")
}

func (c *Container) TemplateName() string {
	return c.templateName
}
