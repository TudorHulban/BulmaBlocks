package container

import (
	"bulma/web"
	"bytes"
	"errors"
	"text/template"
)

// Container Component
type Container struct {
	templateName string

	markdown string
}

var _ web.IWeb = (*Container)(nil)

func NewCo() *Container {
	return &Container{
		templateName: "container.gohtml",
	}
}

func (c *Container) SetMarkdonw(markdown string) {
	c.markdown = markdown
}

func (c *Container) Render(t *template.Template) (string, error) {
	tmpl := t.Lookup(c.templateName)
	if tmpl == nil {
		return "", errors.New("lookup did not work")
	}

	var buf bytes.Buffer

	err := tmpl.ExecuteTemplate(&buf, c.templateName, c)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
