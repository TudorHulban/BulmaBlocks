package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"text/template"

	"bulma/web"
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

	b := web.Breadcumb{
		Categories: []string{"A", "B"},
		Item:       "Tea",
	}

	var buf bytes.Buffer

	s1 := tmpl.Lookup("body.gohtml")
	s1.ExecuteTemplate(&buf, "body", b)

	data := Data{
		Title: "XXX",
		Body:  buf.String(),
	}

	f, errCreate := os.Create("output.html")
	if errCreate != nil {
		fmt.Println("errCreate: ", errCreate)
	}
	defer f.Close()

	tmpl.ExecuteTemplate(f, "layout.gohtml", data)
}
