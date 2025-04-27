package f

import (
	"html/template"
)

func SubmitButton(label string, ops ...submitElement) Submit {
	field := Submit{label: label}

	for _, opt := range ops {
		opt.applySubmitOption(&field)
	}

	return field
}

type Submit struct {
	base
	label string
}

func (s *Submit) Label() template.HTML {
	return ""
}

func (s *Submit) Input() template.HTML {
	str := `<input type="submit" value="` + s.label + `"`

	if s.disabled {
		str += ` disabled`
	}

	str += `/>`

	return template.HTML(str)
}

func (s *Submit) Full() template.HTML {
	return s.Input()
}

func (s *Submit) validate() bool {
	return true
}

// attributes:
// formaction, formenctype, formmethod, formnovalidate, formtarget, accesskey
var (
	_ submitElement = (*disabledOption)(nil)
)
