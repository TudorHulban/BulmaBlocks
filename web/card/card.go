package card

import (
	"bulma/web"
	"bulma/web/image"
)

type CardImage struct {
	image.Content
}

type CardThumbnailImage struct {
	image.Content
}

type Content struct {
	CardImage
	CardThumbnailImage

	Title    string
	SubTitle string
	Text     string
}

// Card Component
type Card struct {
	TemplateName string

	Content
}

var _ web.IWeb = (*Card)(nil)

func NewCo(c Content) *Card {
	return &Card{
		TemplateName: "card.gohtml",

		Content: c,
	}
}

func NewCoTest(c Content) (*Card, error) {
	img1, errImg1 := image.NewImageFixedSize(128, image.Content{
		ImageSrc: "https://bulma.io/images/placeholders/256x256.png",
		ImageAlt: "Image Placeholder",
	})
	if errImg1 != nil {
		return nil, errImg1
	}

	img2, errImg2 := image.NewImageFixedSize(32, image.Content{
		ImageSrc: "https://bulma.io/images/placeholders/32x32.png",
		ImageAlt: "Image Placeholder",
	})

	return &Card{
		TemplateName: "card.gohtml",

		Content: c,
	}, nil
}

func (c *Card) GetTemplateName() string {
	return c.TemplateName
}
