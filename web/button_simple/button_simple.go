package buttonsimple

import (
	"bytes"
	"text/template"
)

// ButtonSimple Component
type ButtonSimple struct {
	templateName string

	Label string
}

// var _ IWeb = (*ButtonSimple)(nil)

// Template Maybe a constructor should be used.
func (c *ButtonSimple) Template() string {
	c.templateName = "button_simple.gohtml"

	return c.templateName
}

func (c *ButtonSimple) Render(t *template.Template) string {
	var buf bytes.Buffer

	t.Lookup(c.templateName).ExecuteTemplate(&buf, "button_simple", c)

	return buf.String()
}
