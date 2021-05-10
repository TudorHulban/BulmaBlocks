package breadcumb

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

func TestBreadcumbFullPage(t *testing.T) {
	cache, errNew := cachetemplates.NewCacher("../../templates")
	require.Nil(t, errNew)

	page, errNew := page.NewLandingPage("Landing Page", cache, page.Content{
		Title: "Landing Page",
	})
	require.Nil(t, errNew)

	// Breadcumb Component
	bread := NewCo("Vegan Burger", cache, Content{
		ActiveItem: "Vegan Burger",
		Categories: []string{"Home", "Categ A", token},
	})

	// bringing now the Body
	body, errNewBody := body.NewCo(cache)
	require.Nil(t, errNewBody)

	require.Nil(t, web.RenderComponentTo(bread, body, cache))

	page.AppendToBody(body.Markdown...)
	page.RenderTo(os.Stdout)
}
