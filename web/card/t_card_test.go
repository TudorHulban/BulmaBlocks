package card

import (
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestCard(t *testing.T) {
	tmpl := template.New("views")

	c := NewCo(Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   token,
		Title:      "Card Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	})

	tmpl, err := tmpl.ParseFiles("../../templates/" + c.TemplateName)
	require.Nil(t, err)

	s, errRender := c.Render(tmpl)
	require.Nil(t, errRender, "Did not render correctly.")
	require.Contains(t, s, token, "Does not contain token.")

	fmt.Println(s)
}
