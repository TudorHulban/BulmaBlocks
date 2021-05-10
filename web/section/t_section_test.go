package section

import (
	"bulma/cachetemplates"
	"bulma/web"
	"bulma/web/body"
	"bulma/web/breadcumb"
	"bulma/web/page"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestContainerFullPage(t *testing.T) {
	cache, errNew := cachetemplates.NewCacher("../../templates")
	require.Nil(t, errNew)

	p, errNew := page.NewLandingPage("Landing Page", cache, page.Content{
		Title: "Landing Page",
	})
	require.Nil(t, errNew)

	// Breadcumb Component
	bread := breadcumb.NewCo("Breadcumb", cache, breadcumb.Content{
		ActiveItem: "Vegan Burger",
		Categories: []string{"Home", "Categ A", token},
	})

	// bringing now the container
	section := NewCo("Section A", cache)

	require.Nil(t, web.RenderComponentTo(bread, section, cache))

	// bringing now the Body
	body := body.NewCo(cache)

	require.Nil(t, web.RenderComponentTo(section, body, cache))

	p.AppendToBody(body.Markdown...)
	p.RenderTo(os.Stdout)
}
