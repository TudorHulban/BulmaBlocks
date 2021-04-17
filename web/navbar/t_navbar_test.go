package navbar

import (
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestContainer(t *testing.T) {
	tmpl := template.New("views")

	tmpl, err := tmpl.ParseFiles("../../templates/navbar.gohtml")
	require.Nil(t, err)

	c := NewCo(Content{
		ItemsNoSubMenu:    []string{"Menu1", "Menu2", token},
		ItemsWithSubMenus: []MenuEntry{},
	})

	s, errRender := c.Render(tmpl)
	require.Nil(t, errRender)
	require.Contains(t, s, token)
}
