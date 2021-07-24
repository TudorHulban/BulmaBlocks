package form

import (
	"bulma/web"

	"github.com/pkg/errors"
)

type FormField struct {
	Label           string
	Placeholder     string
	Value           string
	FieldIconFA     string // font awesome icon name
	CheckIconFA     string // font awesome check icon name - sits on right side
	Hint            string
	Type            string
	SubmitBtnLabel  string
	CancelBtnLabel  string
	DropdownOptions []string
}

type FieldClass string

var FieldClassEnum = map[FieldClass]bool{
	"field":    true,
	"input":    true,
	"control":  true,
	"textarea": true,
	"checkbox": true,
	"radio":    true,
}

// Form Component
type Form struct {
	templateName string

	Item       string
	Categories []string
}

var _ web.IRenderer = (*Form)(nil)

var (
	fieldTemplatePath   = "form_field_simple.gohtml"
	inputTemplatePath   = "form_field_input.gohtml"
	controlTemplatePath = "control.gohtml"
)

func NewCo(i string, cat []string) *Breadcumb {
	return &Breadcumb{
		templateName: "breadcrumb.gohtml",

		Item:       i,
		Categories: cat,
	}
}

func (c *Form) GetTemplateName() string {
	return c.templateName
}

func generateField(field FormField) (string, error) {
	if validateFieldType(field.Type) != nil {
		return "", errors.WithMessage(errValid, "form generation failed due to field validation")
	}

	switch field.Type {
	case "input":
		{
			return inputTemplatePath, nil
		}
	case "control":
		{
			return inputTemplatePath, nil
		}
	}

	return fieldTemplatePath, nil
}

func validateFieldType(f string) error {
	if _, exists := FieldClassEnum[FieldClass(f)]; exists {
		return nil
	}

	return errors.New("field type not known")
}
