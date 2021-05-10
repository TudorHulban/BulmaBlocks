package page

import (
	"bulma/cachetemplates"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPage(t *testing.T) {
	cache, errNew := cachetemplates.NewCacher("../../templates/")
	require.Nil(t, errNew)

	page, errNew := NewLandingPage("Landing Page", cache, Content{
		Title: "Landing Page",
	})
	require.Nil(t, errNew)
	require.NotNil(t, page.templateHTML)

	page.AppendToBody("xxx", "yyy")

	page.RenderTo(os.Stdout)
}
