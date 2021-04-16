package compolayout

// Layout Component
type Layout struct {
	templateName string

	Title string
	Body  string
}

// var _ web.IWeb = (*Layout)(nil)

func NewCo(title, body string) *Layout {
	return &Layout{
		templateName: "body.gohtml",

		Title: title,
		Body:  body,
	}
}
