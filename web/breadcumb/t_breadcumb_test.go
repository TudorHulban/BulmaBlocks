package breadcumb

import (
	"bulma/cachetemplates"
	"bulma/page"
	"bulma/web/body"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestBreadcumbFullPage(t *testing.T) {
	cache, errNew := cachetemplates.NewCacher("../../templates")
	require.Nil(t, errNew)

	p, errNew := page.NewLandingPage("Landing Page", cache, page.Content{
		Title: "Landing Page",
	})
	require.Nil(t, errNew)

	// Breadcumb Component
	c, errNewBreadcumb := NewCo("Breadcumb", cache, Content{
		ActiveItem: "Vegan Burger",
		Categories: []string{"Home", "Categ A", token},
	})
	require.Nil(t, errNewBreadcumb)

	// bringing now the Body
	b, errNewBody := body.NewCo(cache)
	require.Nil(t, errNewBody)

	c.RenderTo(b)
	p.AppendToBody(b.Markdown...)
	p.RenderTo(os.Stdout)
}
