package web

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
)

func Render(t *template.Template, component IWeb) (string, error) {
	tmpl := t.Lookup(component.TemplateName())
	if tmpl == nil {
		return "", errors.Errorf("lookup did not work for template %s", component.TemplateName())
	}

	var buf bytes.Buffer

	err := tmpl.ExecuteTemplate(&buf, component.TemplateName(), component)
	if err != nil {
		return "", err
	}

	return "\n" + buf.String() + "\n", nil
}
