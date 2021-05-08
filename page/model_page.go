package page

import (
	"io"
	"text/template"
)

type Content struct {
	Author string
}

type WebPage struct {
	Name           string
	LayoutTemplate string // to move to layout interface

	Content
}

func NewLandingPage(name string, content Content) (*WebPage, error) {
	return &WebPage{
		Name:           name,
		LayoutTemplate: "hi {{.Content.Author}} to your landing page.",

		Content: content,
	}, nil
}

func (p *WebPage) RenderTo(w io.Writer) error {
	t := template.New(p.Name)

	t, errParse := t.Parse(p.LayoutTemplate)
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, p)
}
