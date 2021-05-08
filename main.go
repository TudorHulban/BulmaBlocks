package main

import (
	"bulma/cachetemplates"
	"bulma/page"
	"bulma/web/body"
	"bytes"

	"net/http"
)

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":7000", nil)
}

func serve(w http.ResponseWriter, r *http.Request) {
	resp, err := prepareContent()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte(resp))
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
