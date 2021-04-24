package mediaobject

import (
	"bulma/web"
)

type Content struct {
	FullName string
	Age      string
	Email    string
	Details  string
}

// Media Component
type Media struct {
	templateName string

	Content
}

var _ web.IWeb = (*Media)(nil)

func NewCo(c Content) *Media {
	return &Media{
		templateName: "media.gohtml",

		Content: c,
	}
}

// func (c *Media) Render(t *template.Template) (string, error) {
// 	return web.Render(t, c.templateName, c)
// }

func (c *Media) GetTemplateName() string {
	return c.templateName
}
