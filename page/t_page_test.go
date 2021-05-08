package page

import (
	"bulma/cachetemplates"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPage(t *testing.T) {
	cache, errNew := cachetemplates.NewCacher()
	require.Nil(t, errNew)

	p, errNew := NewLandingPage("Landing Page", cache, Content{
		Title: "Landing Page",
	})
	require.Nil(t, errNew)
	require.NotNil(t, p.templateHTML)

	p.RenderTo(os.Stdout)
}
