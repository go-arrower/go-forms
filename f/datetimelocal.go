package f

import (
	"html/template"
	"strings"
)

func DateTimeLocalField(label string, ops ...FieldOption) *DateTimeLocal {
	field := &DateTimeLocal{
		base: base{
			label: label,
		},
	}

	for _, opt := range ops {
		opt(field)
	}

	return field
}

type DateTimeLocal struct {
	base
	value      string
	validators []func(string) error
	errors     []Error
	required   bool
}

func (s *DateTimeLocal) ID() string {
	return strings.ToLower(s.label)
}

func (s *DateTimeLocal) Name() string {
	return strings.ToLower(s.label)
}

func (s *DateTimeLocal) Label() template.HTML {
	str := `<label for="` + strings.ToLower(s.label) + `">`
	str += s.label
	str += `</label>`

	return template.HTML(str)
}

func (s *DateTimeLocal) Input() template.HTML {
	str := `<input id="` + strings.ToLower(s.label) + `"`
	str += ` type="datetime-local"`
	str += `name="` + strings.ToLower(s.label) + `" />`

	return template.HTML(str)
}

func (s *DateTimeLocal) Full() template.HTML {
	return s.Label() + s.Input()
}

func (s *DateTimeLocal) Value() string {
	return s.value
}

func (s *DateTimeLocal) SetValue(val string) {
	s.value = val
}

// Validate runs all validators on the field
func (s *DateTimeLocal) Validate() bool {
	for _, validator := range s.validators {
		if err := validator(s.value); err != nil {
			s.errors = append(s.errors, Error{Key: s.label, Message: err.Error()})
			return false
		}
	}

	return true
}
