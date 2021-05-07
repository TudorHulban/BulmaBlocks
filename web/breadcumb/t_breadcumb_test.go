package breadcumb

import (
	"bulma/web"
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestBreadcumb(t *testing.T) {
	c := NewCo("Vegan Burger", []string{"Home", "Categ A", token})

	tmpl := template.New("views")

	tmpl, err := tmpl.ParseFiles(c.getTemplatePath())
	require.Nil(t, err, "Did not parse file "+c.getTemplatePath())

	s, errRender := web.Render(tmpl, c)
	require.Nil(t, errRender, "Did not render correctly.")
	require.Contains(t, s, token, "Does not contain token.")

	fmt.Println(s)
}
