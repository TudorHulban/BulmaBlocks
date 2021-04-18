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
	"bulma/web/navbar"
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

	nav := navbar.NewCo(navbar.Content{
		ItemsNoSubMenu: []string{"Menu1", "Menu2", "Menu3"},
		ItemsWithSubMenus: []navbar.MenuEntry{
			navbar.MenuEntry{Menu: "XXX",
				Entries: []string{"XXX-A", "XXX-B", "XXX-C"},
			},
			navbar.MenuEntry{Menu: "YYY",
				Entries: []string{"YYY-A"},
			},
		},
	})

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
	errInject := body.Inject(tmpl, c, nav, b, m1, m2)
	if errInject != nil {
		fmt.Println(errInject)
		os.Exit(2)
	}

	f, errCreate := os.Create("output.html")
	if errCreate != nil {
		fmt.Println("errCreate: ", errCreate)
	}
	defer f.Close()

	l := layout.NewCo("This is title", body.Markdown())
	errExe := tmpl.ExecuteTemplate(f, "layout.gohtml", l)
	if errExe != nil {
		fmt.Println(errExe)
		os.Exit(3)
	}
}
