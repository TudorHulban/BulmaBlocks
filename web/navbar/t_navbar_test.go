package navbar

import (
	"bulma/web"
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestNavbar(t *testing.T) {
	c := NewCo(Content{
		ItemsNoSubMenu: []string{"Menu1", "Menu2", token},
		ItemsWithSubMenus: []MenuEntry{MenuEntry{
			Menu:    "XXX",
			Entries: []string{"XXX-A", "XXX-B", "XXX-C"},
		}},
	})

	tmpl := template.New("views")

	tmpl, err := tmpl.ParseFiles(c.GetTemplateName())
	require.Nil(t, err)

	s, errRender := web.Render(tmpl, c)
	require.Nil(t, errRender, "Did not render correctly.")
	require.Contains(t, s, token, "Does not contain token.")

	fmt.Println(s)
}
