package main

import (
	"bulma/cachetemplates"
	"bulma/page"
	"bulma/web/body"
	"bytes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		resp, err := prepareContent()
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.SendString(string(resp))
	})

	app.Listen(":7000")
}

func prepareContent() (cachetemplates.HTML, error) {
	cache, errCache := cachetemplates.NewCacher("./templates")
	if errCache != nil {
		return nil, errCache
	}

	p, errPage := page.NewLandingPage("Landing Page", cache, page.Content{
		Title: "Landing Page",
	})
	if errPage != nil {
		return nil, errPage
	}

	b, errBody := body.NewCo(cache)
	if errBody != nil {
		return nil, errBody
	}

	b.AppendToBody("xx1", "yy2")
	p.AppendToBody(b.Markdown...)

	var res bytes.Buffer
	errRender := p.RenderTo(&res)

	return cachetemplates.HTML(res.String()), errRender
}
