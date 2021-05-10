package image

import (
	"bulma/cachetemplates"
	"bulma/web"
	"errors"
	"strings"
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
	ImageURL             string
	ImageAlt             string
	AdditionalCSSClasses []string
}

// Image Component
type Image struct {
	Name         string
	templateName string

	CSSClass string

	Content
}

var _ web.IRenderer = (*Image)(nil)

func validateFixedSize(square int) bool {
	_, exists := fixedSquareSizes[square]

	return exists
}

func NewImageFixedSize(componentName string, size int, templates map[cachetemplates.TemplatePath]cachetemplates.HTML, content Content) (*Image, error) {
	if !validateFixedSize(size) {
		return nil, errors.New("passed image size is not predefined")
	}

	res := Image{
		Name:         componentName,
		templateName: "image.gohtml",

		Content: content,
	}

	res.IsImageSize = res.isImageSizeFixed(size)

	return &res, nil
}

func NewImageRounded(componentName string, templates map[cachetemplates.TemplatePath]cachetemplates.HTML, content Content) *Image {
	return &Image{
		Name:         componentName,
		templateName: "image.gohtml",
		CSSClass:     "is-rounded",

		Content: content,
	}
}

func (i *Image) TemplateName() string {
	return i.templateName
}

func (i *Image) prepareCSS() {
	if len(i.AdditionalCSSClasses) == 0 {
		return
	}

	i.AdditionalCSSClasses = append(i.AdditionalCSSClasses, i.CSSClass)

	i.CSSClass = strings.Join(i.AdditionalCSSClasses, " ")
}

func (*Image) isImageSizeFixed(square int) string {
	v, exists := fixedSquareSizes[square]
	if !exists {
		return ""
	}

	return strings.Join([]string{"is", "-", v, "x", v}, "")
}
