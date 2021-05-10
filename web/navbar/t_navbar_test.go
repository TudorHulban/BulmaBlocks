package navbar

import (
	"bulma/cachetemplates"
	"bulma/web"
	"bulma/web/body"
	"bulma/web/page"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestNavbar(t *testing.T) {
	cache, errNew := cachetemplates.NewCacher("../../templates")
	require.Nil(t, errNew)

	page, errNew := page.NewLandingPage("Landing Page", cache, page.Content{
		Title: "Landing Page",
	})
	require.Nil(t, errNew)

	nav := NewCo("Navigation", cache, Content{
		ItemsNoSubMenu: []string{"Menu1", "Menu2", token},
		ItemsWithSubMenus: []MenuEntry{MenuEntry{
			Menu:    "XXX",
			Entries: []string{"XXX-A", "XXX-B", "XXX-C"},
		}},
	})

	// bringing now the Body
	body, errNewBody := body.NewCo(cache)
	require.Nil(t, errNewBody)

	require.Nil(t, web.RenderComponentTo(nav, body, cache))

	page.AppendToBody(body.Markdown...)
	page.RenderTo(os.Stdout)
}
