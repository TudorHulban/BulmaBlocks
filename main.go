package main

import (
	"bulma/cachetemplates"
	"bulma/web"
	"bulma/web/body"
	"bulma/web/breadcumb"
	"bulma/web/navbar"
	"bulma/web/page"
	"bulma/web/section"
	"bytes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", serve)

	log.Fatal(app.Listen(":7000"))
}

func serve(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK).Response().Header.SetContentType("text/html")

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

	bread, errNewBreadcumb := breadcumb.NewCo("Breadcumb", cache, breadcumb.Content{
		ActiveItem: "Vegan Burger",
		Categories: []string{"Home", "Categ A", "Categ B"},
	})
	if errNewBreadcumb != nil {
		return nil, errNewBreadcumb
	}

	// bringing now the container
	section, errNewContainer := section.NewCo("Container", cache)
	if errNewContainer != nil {
		return nil, errNewContainer
	}

	body, errBody := body.NewCo(cache)
	if errBody != nil {
		return nil, errBody
	}

	// TODO: check errors
	web.RenderComponentTo(nav, body, cache)
	web.RenderComponentTo(bread, section, cache)
	web.RenderComponentTo(section, body, cache)

	page.AppendToBody(body.Markdown...)

	var res bytes.Buffer
	errRender := page.RenderTo(&res)

	return cachetemplates.HTML(res.String()), errRender
}
