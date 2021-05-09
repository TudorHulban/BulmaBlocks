package webcontainers

import (
	"bulma/web"
	"text/template"

	"github.com/pkg/errors"
)

// Inject takes several components which it renders and decorates with the passed container.
func Inject(t *template.Template, container IWebContainer, blocks ...web.IWeb) error {
	var content []string

	for i, block := range blocks {
		markdown, err := web.Render(t, block) //block.Render(t)
		if err != nil {
			return errors.WithMessagef(err, "at injection of component number %v", i)
		}

		content = append(content, markdown)
	}

	container.SetContent(content)
	return nil
}
