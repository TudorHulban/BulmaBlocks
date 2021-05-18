package main

import (
	"bulma/cachetemplates"
	"bulma/web"
	"bulma/web/body"
	"bulma/web/breadcumb"
	"bulma/web/image"
	"bulma/web/navbar"
	"bulma/web/page"
	"bulma/web/section"
	"bytes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public")
	app.Get("/", serve)

	log.Fatal(app.Listen(":7000"))
}

func serve(c *fiber.Ctx) error {
	resp, err := prepareContent()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Response().SetBody([]byte(err.Error()))
	}

	c.Type("html").Send([]byte(resp))
	return nil
}

func prepareContent() (cachetemplates.HTML, error) {
	cache, errCache := cachetemplates.NewCacher("./templates")
	if errCache != nil {
		return nil, errCache
	}

	page, errPage := page.NewLandingPage("Landing Page", cache, page.Content{
		Title: "Landing Page",
	})
	if errPage != nil {
		return nil, errPage
	}

	nav := navbar.NewCo("Navigation", cache, navbar.Content{
		ItemsNoSubMenu: []string{"Menu1", "Menu2", "Menu3"},
		ItemsWithSubMenus: []navbar.MenuEntry{navbar.MenuEntry{
			Menu:    "XXX",
			Entries: []string{"XXX-A", "XXX-B", "XXX-C"},
		}},
	})

	bread := breadcumb.NewCo("Breadcumb", cache, breadcumb.Content{
		ActiveItem: "Vegan Burger",
		Categories: []string{"Home", "Categ A", "Categ B"},
	})

	img := image.NewImageRounded("Round Image", cache, image.Content{
		ImageSrc: "pexels_1.jpg",
		ImageAlt: "rounded pexel image",
	})

	// bringing now the container
	section := section.NewCo("Container", cache)

	body := body.NewCo(cache)

	// TODO: check errors
	web.RenderComponentTo(nav, body, cache)
	web.RenderComponentTo(bread, section, cache)
	web.RenderComponentTo(img, section, cache)
	web.RenderComponentTo(section, body, cache)

	page.AppendToBody(body.Markdown...)

	var res bytes.Buffer
	errRender := page.RenderTo(&res)

	return cachetemplates.HTML(res.String()), errRender
}
