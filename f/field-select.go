package f

import (
	"html/template"
	"strings"
)

func SelectField(label string, choices map[string]string, ops ...FieldOption) *Select {
	field := &Select{
		base: base{
			label: label,
		},
		choices: choices,
	}

	// for _, opt := range ops {
	// 	opt(field)
	// }

	return field
}

type Select struct {
	base
	choices    map[string]string
	value      string
	validators []func(string) error
	errors     []Error
	required   bool
}

func (s *Select) ID() string {
	return strings.ToLower(s.label)
}

func (s *Select) Name() string {
	return strings.ToLower(s.label)
}

func (s *Select) Label() template.HTML {
	str := `<label for="` + strings.ToLower(s.label) + `">`

	str += s.label
	if s.required {
		str += ` *`
	}

	str += `</label>`

	return template.HTML(str)
}

func (s *Select) Input() template.HTML {
	str := `<select id="` + strings.ToLower(s.label) + `"`
	str += ` name="` + strings.ToLower(s.label) + `">`

	for value, name := range s.choices {
		str += `<option value="` + value + `">` + name + `</option>`
	}

	str += `</select>`

	return template.HTML(str)
}

func (s *Select) Full() template.HTML {
	return s.Label() + s.Input()
}

func (s *Select) Value() string {
	return s.value
}

func (s *Select) SetValue(val string) {
	s.value = val
}

// Validate runs all validators on the inputElement
func (s *Select) Validate() bool {
	for _, validator := range s.validators {
		if err := validator(s.value); err != nil {
			s.errors = append(s.errors, Error{Key: s.label, Message: err.Error()})
			return false
		}
	}

	return true
}
