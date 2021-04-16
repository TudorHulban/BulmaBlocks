package web

import (
	"bytes"
	"text/template"
)

// Body Component
type Body struct {
	TemplateName string

	Title    string
	Markdown string
}

var _ IWeb = (*Body)(nil)

// Template Maybe a constructor should be used.
func (c *Body) Template() string {
	c.TemplateName = "body.gohtml"

	return c.TemplateName
}

func (c *Body) Render(t *template.Template, data ...interface{}) string {
	var buf bytes.Buffer

	t.Lookup(c.TemplateName).ExecuteTemplate(&buf, "body", data)

	return buf.String()
}
