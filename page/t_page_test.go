package page

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPage(t *testing.T) {
	p, errNew := NewLandingPage("Landing Page", Content{
		Author: "John",
	})
	require.Nil(t, errNew)

	p.RenderTo(os.Stdout)
}
