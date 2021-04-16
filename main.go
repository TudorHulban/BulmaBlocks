package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"bulma/web/compobody"
	"bulma/web/compobreadcumb"
)

type Data struct {
	Title string
	Body  string
}

//go:embed templates/*.gohtml
var f embed.FS

func main() {
	tmpl := template.New("views")

	tmpl, err := tmpl.ParseFS(f, "templates/*.gohtml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b := compobreadcumb.NewCo("Tea", []string{"A", "B"})
	s, err := b.Render(tmpl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body := compobody.Body{}
	body.AddMarkdown(s)

	data := body.Markdown()

	f, errCreate := os.Create("output.html")
	if errCreate != nil {
		fmt.Println("errCreate: ", errCreate)
	}
	defer f.Close()

	tmpl.ExecuteTemplate(f, "layout.gohtml", data)
}
