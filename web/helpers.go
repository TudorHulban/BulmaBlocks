package web

import (
	"bytes"
	"errors"
	"text/template"
)

func Render(t *template.Template, component IWeb) (string, error) {
	tmpl := t.Lookup(component.GetTemplateName())
	if tmpl == nil {
		return "", errors.New("lookup did not work")
	}

	var buf bytes.Buffer

	err := tmpl.ExecuteTemplate(&buf, component.GetTemplateName(), component)
	if err != nil {
		return "", err
	}

	return "\n" + buf.String() + "\n", nil
}
