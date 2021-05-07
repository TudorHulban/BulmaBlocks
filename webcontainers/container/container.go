package container

import (
	"bulma/web"
	"bulma/webcontainers"
	"strings"
)

type Content struct {
	Markdown string
}

// Container Component
type Container struct {
	templateName string

	Content
}

const goTemplate = "container.gohtml"

var _ webcontainers.IWebContainer = (*Container)(nil)

func NewCo(testMode ...bool) *Container {
	production := func() *Container {
		return &Container{
			templateName: goTemplate,
		}
	}

	if testMode == nil {
		return production()
	}

	if testMode[0] == true {
		return &Container{
			templateName: web.TemplateFolderPath + goTemplate,
		}
	}

	return production()
}

// SetContent Method to be used for injecting markdown (ex. from already rendered templates)
func (c *Container) SetContent(content []string) {
	c.Content.Markdown = strings.Join(content, "n")
}

func (c *Container) GetTemplateName() string {
	return c.templateName
}
