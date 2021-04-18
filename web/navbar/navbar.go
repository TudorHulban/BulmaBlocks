package navbar

import (
	"bulma/web"
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
	return web.Render(t, c.templateName, c)
}
