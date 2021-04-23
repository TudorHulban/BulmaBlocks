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

func validateFixedSize(square int) bool {
	_, exists := fixedSquareSizes[square]

	return exists
}

func NewImageFixedSize(size int, c Content) *Image {
	if !validateFixedSize(size) {
		return nil
	}

	res := Image{
		TemplateName: "image.gohtml",

		Content: c,
	}

	res.IsImageSize = res.isImageSizeFixed(size)

	return &res
}

func NewImageRounded(c Content) *Image {
	return &Image{
		TemplateName: "image.gohtml",
		CSSClass:     "is-rounded",

		Content: c,
	}
}

func (c *Image) Render(t *template.Template) (string, error) {
	c.prepareCSS()

	return web.Render(t, c.TemplateName, c)
}

func (c *Image) prepareCSS() {
	if len(c.AdditionalCSSClasses) == 0 {
		return
	}

	c.AdditionalCSSClasses = append(c.AdditionalCSSClasses, c.CSSClass)

	c.CSSClass = strings.Join(c.AdditionalCSSClasses, " ")
}

func (c *Image) isImageSizeFixed(square int) string {
	v, exists := fixedSquareSizes[square]
	if !exists {
		return ""
	}

	return strings.Join([]string{"is", "-", v, "x", v}, "")
}