package breadcumb

import (
	"bulma/cachetemplates"
	"bulma/page"
	"bulma/web"
	"bulma/web/body"
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
	bread, errNewBreadcumb := NewCo("Vegan Burger", cache, Content{
		ActiveItem: "Vegan Burger",
		Categories: []string{"Home", "Categ A", token},
	})
	require.Nil(t, errNewBreadcumb)

	// bringing now the Body
	body, errNewBody := body.NewCo(cache)
	require.Nil(t, errNewBody)

	web.RenderComponentTo(bread, body, cache)

	page.AppendToBody(body.Markdown...)
	page.RenderTo(os.Stdout)
}
