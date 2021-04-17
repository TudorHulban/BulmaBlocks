package container

import (
	"embed"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

//go:embed container.gohtml
var f embed.FS

func TestContainer(t *testing.T) {
	tmpl := template.New("views")

	tmpl, err := tmpl.ParseFS(f, "container.gohtml")
	require.Nil(t, err)

	c := NewCo()
	c.SetMarkdown(token)

	s, errRender := c.Render(tmpl)
	require.Nil(t, errRender)
	require.Contains(t, s, token)
}
