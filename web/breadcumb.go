package web

// Breadcumb Component
type Breadcumb struct {
	TemplateName string

	Item       string
	Categories []string
}

var _ IWeb = (*Breadcumb)(nil)

// Template Maybe a constructor should be used.
func (c *Breadcumb) Template() string {
	c.TemplateName = "body.gohtml"

	return c.TemplateName
}
