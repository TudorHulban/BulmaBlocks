package breadcumb

import (
	"bulma/web"
	"bulma/web/body"

	"bulma/webcontainers/container"
	"bytes"
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

func TestBreadcumbFullPage(t *testing.T) {
	c := NewCo("Vegan Burger", []string{"Home", "Categ A", token})

	tmplBreadcumb := template.New(c.templateName)
	tmplBreadcumb, err := tmplBreadcumb.ParseFiles(c.getTemplatePath())
	require.Nil(t, err, "Did not parse file "+c.getTemplatePath())

	s, errRender := web.Render(tmplBreadcumb, c)
	require.Nil(t, errRender, "Did not render correctly.")
	require.Contains(t, s, token, "Does not contain token.")

	co := container.NewCo(true)
	co.SetContent([]string{s})

	tmpl := template.New("views")

	tmpl, err := tmpl.ParseFiles(c.getTemplatePath())
	require.Nil(t, err, "Did not parse file "+c.getTemplatePath())

	body := compobody.Body{}

	var buf bytes.Buffer

	errExe := tmpl.ExecuteTemplate(&buf, body.GetTemplateName(), body)
	require.Nil(t, errExe)

	// errInject := body.Inject(tmpl, co, c)
	// require.Nil(t, errInject, "Component could not be injected in body")

	// s, errRender := web.Render(tmpl, c)
	// require.Nil(t, errRender, "Did not render correctly.")
	// require.Contains(t, s, token, "Does not contain token.")

	// fmt.Println(s)
}
