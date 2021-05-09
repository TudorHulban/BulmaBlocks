package container

import (
	"bulma/cachetemplates"
	"bulma/page"
	"bulma/web/body"
	"bulma/web/breadcumb"
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
	bread, errNewBreadcumb := breadcumb.NewCo("Breadcumb", cache, breadcumb.Content{
		ActiveItem: "Vegan Burger",
		Categories: []string{"Home", "Categ A", token},
	})
	require.Nil(t, errNewBreadcumb)

	// bringing now the container
	container, errNewContainer := NewCo("Container", cache)
	require.Nil(t, errNewContainer)

	bread.RenderTo(container)

	// bringing now the Body
	b, errNewBody := body.NewCo(cache)
	require.Nil(t, errNewBody)

	container.RenderTo(b)

	p.AppendToBody(b.Markdown...)
	p.RenderTo(os.Stdout)
}
