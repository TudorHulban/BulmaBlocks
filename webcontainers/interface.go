package webcontainers

import (
	"text/template"
)

type IWebContainer interface {
	SetMarkdown(markdown string)
	Render(t *template.Template) (string, error)
}
