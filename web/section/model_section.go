package section

import (
	"bulma/cachetemplates"
	"bulma/web"
)

// Section Component
type Section struct {
	Name         string
	templateName string

	Markdown []string
}

var _ web.IRenderer = (*Section)(nil)

func NewCo(sectionName string, templates map[cachetemplates.TemplatePath]cachetemplates.HTML) *Section {
	return &Section{
		Name:         sectionName,
		templateName: "section.gohtml",
	}
}

func (c *Section) TemplateName() string {
	return c.templateName
}

func (c *Section) Write(markdown []byte) (int, error) {
	c.Markdown = append(c.Markdown, string(markdown))

	return 0, nil
}
