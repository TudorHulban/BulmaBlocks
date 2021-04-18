package webcontainers

import (
	"bulma/web"
	"text/template"
)

func Render(t *template.Template, compoTemplateName string, component web.IWeb) (string, error) {
	return web.Render(t, compoTemplateName, component)
}
