package navbar

import (
	"bulma/web"
	"bytes"
	"errors"
	"text/template"
)

type MenuEntry struct {
	Menu    string
	Entries []string
}

type Content struct {
	LogoLink          string
	ItemsNoSubMenu    []string
	ItemsWithSubMenus []MenuEntry
}

// Layout Component
type Navbar struct {
	templateName string

	Content
}

var _ web.IWeb = (*Navbar)(nil)

func NewCo(c Content) *Navbar {
	return &Navbar{
		templateName: "navbar.gohtml",

		Content: c,
	}
}

func (c *Navbar) Render(t *template.Template) (string, error) {
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
