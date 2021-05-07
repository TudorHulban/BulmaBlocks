package container

import (
	"bulma/webcontainers"
	// "bytes"
	// "errors"
	// "text/template"
)

// Container Component
type Container struct {
	templateName string

	Markdown string
}

var _ webcontainers.IWebContainer = (*Container)(nil)

func NewCo() *Container {
	return &Container{
		templateName: "container.gohtml",
	}
}

func (c *Container) SetMarkdown(markdown string) {
	c.Markdown = markdown
}

func (c *Container) GetTemplateName() string {
	return c.templateName
}
