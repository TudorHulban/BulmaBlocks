package card

import (
	"bulma/web/image"
	"fmt"
	"testing"
	// "text/template"
	// "github.com/stretchr/testify/require"
)

const token = "xxx"

func TestCard(t *testing.T) {
	// tmpl := template.New("views")

	img1 := image.NewImageFixedSize(256, image.Content{
		ImageSrc: "https://bulma.io/images/placeholders/256x256.png",
		ImageAlt: token,
	})

	img2 := image.NewImageFixedSize(32, image.Content{
		ImageSrc: "https://bulma.io/images/placeholders/32x32.png",
		ImageAlt: token,
	})

	card := NewCo(Content{
		CardImage: CardImage{
			img1.Content,
		},
		CardThumbnailImage: CardThumbnailImage{
			img2.Content,
		},

		Title:    "This is card title",
		SubTitle: "Subtitle",
		Text:     "Lorem",
	})

	fmt.Println(card)

	// tmpl, err := tmpl.ParseFiles("../../templates/" + card.TemplateName)
	// require.Nil(t, err)

	// s, errRender := card.Render(tmpl)
	// require.Nil(t, errRender, "Did not render correctly.")
	// require.Contains(t, s, token, "Does not contain token.")

	// fmt.Println(s)
}
