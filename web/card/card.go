package card

import (
	"bulma/web"
	"bulma/web/image"
	"text/template"
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

func (c *Card) Render(t *template.Template) (string, error) {
	return web.Render(t, c.TemplateName, c)
}
