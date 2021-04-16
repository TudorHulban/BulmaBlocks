package compobreadcumb

import (
	"bytes"
	"errors"
	"text/template"
)

// Breadcumb Component
type Breadcumb struct {
	TemplateName string

	Item       string
	Categories []string
}

// var _ web.IWeb = (*Breadcumb)(nil)

func NewCo(i string, cat []string) *Breadcumb {
	return &Breadcumb{
		TemplateName: "breadcrumb.gohtml",
		Item:         i,
		Categories:   cat,
	}
}

func (c *Breadcumb) Render(t *template.Template) (string, error) {
	tmpl := t.Lookup(c.TemplateName)
	if tmpl == nil {
		return "", errors.New("lookup did not work")
	}

	var buf bytes.Buffer

	err := tmpl.ExecuteTemplate(&buf, c.TemplateName, c)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
