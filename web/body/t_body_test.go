package body

import (
	"bulma/cachetemplates"
	"bulma/page"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBodyAppend(t *testing.T) {
	cache, errNew := cachetemplates.NewCacher("../../templates/")
	require.Nil(t, errNew)

	p, errNew := page.NewLandingPage("Landing Page", cache, page.Content{
		Title: "Landing Page",
	})
	require.Nil(t, errNew)

	// bringing now the body
	b, errNew := NewCo(cache)
	require.Nil(t, errNew)

	b.AppendToBody("xx1", "yy2")
	p.AppendToBody(b.Markdown...)

	p.RenderTo(os.Stdout)
}
