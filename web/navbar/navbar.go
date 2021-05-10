package navbar

import (
	"bulma/cachetemplates"
	"bulma/web"
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
	Name         string
	templateName string
	templateHTML []byte

	Content
}

var _ web.IRenderer = (*Navbar)(nil)

func NewCo(componentName string, templates map[cachetemplates.TemplatePath]cachetemplates.HTML, content Content) *Navbar {
	return &Navbar{
		Name:         componentName,
		templateName: "navbar.gohtml",

		Content: content,
	}
}

func (n *Navbar) TemplateName() string {
	return n.templateName
}
