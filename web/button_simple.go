package web

// ButtonSimple Component
type ButtonSimple struct {
	TemplateName string
}

var _ IWeb = (*ButtonSimple)(nil)

// Template Maybe a constructor should be used.
func (c *ButtonSimple) Template() string {
	c.TemplateName = "button_simple.gohtml"

	return c.TemplateName
}
