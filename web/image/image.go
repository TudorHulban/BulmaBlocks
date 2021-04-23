package image

import (
	"bulma/web"
	"strings"
	"text/template"
)

var fixedSquareSizes = map[int]string{
	16:  "16",
	24:  "24",
	32:  "32",
	48:  "48",
	64:  "64",
	96:  "96",
	128: "128",
}

type Content struct {
	IsImageSize          string
	ImageSrc             web.URL
	ImageAlt             string
	AdditionalCSSClasses []string
}

// Image Component
type Image struct {
	TemplateName string
	CSSClass     string

	Content
}

var _ web.IWeb = (*Image)(nil)

func NewImage(c Content) *Image {
	return &Image{
		TemplateName: "image.gohtml",
		CSSClass:     "img-fluid",

		Content: c,
	}
}

func NewImageThumbnail(c Content) *Image {
	return &Image{
		TemplateName: "image.gohtml",
		CSSClass:     "img-thumbnail",

		Content: c,
	}
}

func (c *Image) Render(t *template.Template) (string, error) {
	c.prepareCSS()

	return web.Render(t, c.TemplateName, c)
}

func (c *Image) prepareCSS() {
	c.AdditionalCSSClasses = append(c.AdditionalCSSClasses, c.CSSClass)

	c.CSSClass = strings.Join(c.AdditionalCSSClasses, " ")
}

func (c *Image) IsImageSizeFixed(square int) string {
	v, exists := fixedSquareSizes[square]
	if !exists {
		return ""
	}

	return strings.Join([]string{"is", "-", v, "x", v}, "")
}
