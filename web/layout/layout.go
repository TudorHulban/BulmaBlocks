package layout

// Layout Component
type Layout struct {
	templateName string

	Title string
	Body  string
}

func NewCo(title, body string) *Layout {
	return &Layout{
		templateName: "body.gohtml",

		Title: title,
		Body:  body,
	}
}
