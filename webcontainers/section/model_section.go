package section

import (
	"bulma/cachetemplates"
	"errors"
	"io"
	"strings"
	"text/template"
)

// Section Component
type Section struct {
	Name         string
	templateName string
	templateHTML []byte

	Markdown []string
}

func NewCo(sectionName string, templates map[cachetemplates.TemplatePath]cachetemplates.HTML) (*Section, error) {
	section := Section{
		Name:         sectionName,
		templateName: "section.gohtml",
	}

	var err error

	section.templateHTML, err = section.getTemplateHTML(templates)
	if err != nil {
		return nil, err
	}

	return &section, nil
}

func (c *Section) Write(markdown []byte) (int, error) {
	c.Markdown = append(c.Markdown, string(markdown))

	return 0, nil
}

func (c *Section) RenderTo(w io.Writer) error {
	t := template.New(c.Name)

	t, errParse := t.Parse(string(c.templateHTML))
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, c)
}

func (c *Section) getTemplateHTML(templates map[cachetemplates.TemplatePath]cachetemplates.HTML) ([]byte, error) {
	for k, html := range templates {
		if strings.Contains(string(k), c.templateName) {
			return html, nil
		}
	}

	return nil, errors.New("no template was found for page")
}

func (c *Section) TemplateName() string {
	return c.templateName
}
