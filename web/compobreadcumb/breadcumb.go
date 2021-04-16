package compobreadcumb

import (
	"bulma/web"
	"bytes"
	"errors"
	"text/template"
)

// Breadcumb Component
type Breadcumb struct {
	templateName string

	Item       string
	Categories []string
}

var _ web.IWeb = (*Breadcumb)(nil)

func NewCo(i string, cat []string) *Breadcumb {
	return &Breadcumb{
		templateName: "breadcumb.gohtml",
		Item:         i,
		Categories:   cat,
	}
}

// Template Maybe a constructor should be used.
func (c *Breadcumb) Template() string {
	c.templateName = "breadcumb.gohtml"

	return c.templateName
}

func (c *Breadcumb) Render(t *template.Template) (string, error) {
	tmpl := t.Lookup(c.templateName)
	if tmpl == nil {
		return "", errors.New("lookup did not work")
	}

	var buf bytes.Buffer

	err := tmpl.ExecuteTemplate(&buf, "breadcumb", *c)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
