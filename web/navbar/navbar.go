package navbar

import (
	"bytes"
	"errors"
	"text/template"
)

type MenuEntry struct {
	Menu    string
	Entries []string
}

type Content struct {
	ItemsNoSubMenu    []string
	ItemsWithSubMenus []MenuEntry
}

// Layout Component
type Navbar struct {
	templateName string

	Content
}

var _ web.IWeb = (*Navbar)(nil)

func NewCo(c Content) *Media {
	return &Media{
		templateName: "media.gohtml",

		Content: c,
	}
}

func (c *Media) Render(t *template.Template) (string, error) {
	tmpl := t.Lookup(c.templateName)
	if tmpl == nil {
		return "", errors.New("lookup did not work")
	}

	var buf bytes.Buffer

	err := tmpl.ExecuteTemplate(&buf, c.templateName, c)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
