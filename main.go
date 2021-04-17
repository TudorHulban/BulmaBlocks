package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"bulma/web/body"
	"bulma/web/breadcumb"
	"bulma/web/layout"
	"bulma/web/media_object"
	"bulma/webcontainers/container"
)

//go:embed templates/*.gohtml
var f embed.FS

func main() {
	tmpl := template.New("views")

	tmpl, err := tmpl.ParseFS(f, "templates/*.gohtml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b := breadcumb.NewCo("Tea", []string{"A", "B"})

	m1 := mediaobject.NewCo(mediaobject.Content{
		FullName: "John Smith",
		Age:      "44",
		Email:    "john@gmx.de",
		Details:  "Life is beautiful.",
	})

	m2 := mediaobject.NewCo(mediaobject.Content{
		FullName: "Maurice Ravel",
		Age:      "74",
		Email:    "",
		Details:  "Music is everything.",
	})

	c := container.NewCo()

	body := compobody.Body{}
	body.Inject(tmpl, c, b, m1, m2)

	f, errCreate := os.Create("output.html")
	if errCreate != nil {
		fmt.Println("errCreate: ", errCreate)
	}
	defer f.Close()

	l := layout.NewCo("This is title", body.Markdown())
	tmpl.ExecuteTemplate(f, "layout.gohtml", l)
}
