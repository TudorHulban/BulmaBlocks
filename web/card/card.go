package card

import (
	"blocks/web"
	"text/template"
)

type Content struct {
	ImageSrc   web.URL
	ImageAlt   string
	Title      string
	Text       string
	ButtonText string
	ButtonURL  web.URL
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
