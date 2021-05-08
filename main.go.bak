package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"bulma/web/body"
	"bulma/web/breadcumb"
	"bulma/web/card"
	"bulma/web/image"
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

	img1, errImg1 := image.NewImageFixedSize(128, image.Content{
		ImageSrc: "https://bulma.io/images/placeholders/256x256.png",
		ImageAlt: "Place Holder",
	})
	if errImg1 != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	img2, errImg2 := image.NewImageFixedSize(32, image.Content{
		ImageSrc: "https://bulma.io/images/placeholders/32x32.png",
		ImageAlt: "Place Holder",
	})
	if errImg2 != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	ccontent := card.Content{
		Title:    "This is card title",
		SubTitle: "Subtitle",
		Text:     "Lorem",
	}

	ccontent.CardImage = card.CardImage{
		img1.Content,
	}
	ccontent.CardThumbnailImage = card.CardThumbnailImage{
		img2.Content,
	}

	card := card.NewCo(ccontent)

	c := container.NewCo()

	body := compobody.Body{}
	errInject := body.Inject(tmpl, c, nav, b, m1, m2, card)
	if errInject != nil {
		fmt.Println(errInject)
		os.Exit(4)
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
		os.Exit(5)
	}
}
