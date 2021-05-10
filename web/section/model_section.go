package section

import (
	"bulma/cachetemplates"
)

// Section Component
type Section struct {
	Name         string
	templateName string
	templateHTML []byte

	Markdown []string
}

var _ web.IRenderer = (*Section)(nil)

func NewCo(sectionName string, templates map[cachetemplates.TemplatePath]cachetemplates.HTML) (*Section, error) {
	section := Section{
		Name:         sectionName,
		templateName: "section.gohtml",
	}

	return &section, nil
}

func (c *Section) Write(markdown []byte) (int, error) {
	c.Markdown = append(c.Markdown, string(markdown))

	return 0, nil
}

func (c *Section) TemplateName() string {
	return c.templateName
}
