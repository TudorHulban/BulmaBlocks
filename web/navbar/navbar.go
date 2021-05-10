package navbar

import (
	"bulma/cachetemplates"
	"bulma/web"
	"errors"
	"io"
	"strings"
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
	Name         string
	templateName string
	templateHTML []byte

	Content
}

var _ web.IRenderer = (*Navbar)(nil)

func NewCo(c Content) *Navbar {
	return &Navbar{
		templateName: "navbar.gohtml",

		Content: c,
	}
}

func (n *Navbar) RenderTo(w io.Writer) error {
	t := template.New(n.Name)

	t, errParse := t.Parse(string(n.templateHTML))
	if errParse != nil {
		return errParse
	}

	return t.Execute(w, n)
}

func (n *Navbar) templatePath() string {
	return web.TemplateFolderPath + n.templateName
}

func (n *Navbar) getTemplateHTML(templates map[cachetemplates.TemplatePath]cachetemplates.HTML) ([]byte, error) {
	for k, html := range templates {
		if strings.Contains(string(k), n.templateName) {
			return html, nil
		}
	}

	return nil, errors.New("no template was found for page")
}
