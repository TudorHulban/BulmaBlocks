package image

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

func TestImageFullPage(t *testing.T) {
	cache, errNew := cachetemplates.NewCacher("../../templates")
	require.Nil(t, errNew)

	page, errNew := page.NewLandingPage("Landing Page", cache, page.Content{
		Title: "Landing Page",
	})
	require.Nil(t, errNew)

	// Image Component
	img := NewImageRounded("Round Image", cache, Content{
		ImageSrc: "../../public/pexels_1.jpg",
	})

	// bringing now the Body
	body := body.NewCo(cache)

	require.Nil(t, web.RenderComponentTo(img, body, cache))

	page.AppendToBody(body.Markdown...)
	page.RenderTo(os.Stdout)
}
