package web

// Layout Component
type Layout struct {
	TemplateName string
}

var _ IWeb = (*Layout)(nil)

// Template Maybe a constructor should be used.
func (c *Layout) Template() string {
	c.TemplateName = "layout.gohtml"

	return c.TemplateName
}
