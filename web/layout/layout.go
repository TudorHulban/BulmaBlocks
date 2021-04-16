package web

import (
	"bulma/web"
)

// Layout Component
type Layout struct {
	templateName string
}

var _ web.IWeb = (*Layout)(nil)

// Template Maybe a constructor should be used.
func (c *Layout) Template() string {
	c.templateName = "layout.gohtml"

	return c.templateName
}
