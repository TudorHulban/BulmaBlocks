package body

import (
	"bulma/cachetemplates"
	"bulma/web/page"
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

	body := NewCo(cache)

	body.AppendToBody("xx1", "yy2")
	p.AppendToBody(body.Markdown...)

	p.RenderTo(os.Stdout)
}
