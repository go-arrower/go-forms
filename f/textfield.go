package f

import (
	"html/template"
	"strings"
)

// TextField constructs a new Text form field.
func TextField(label string, ops ...FieldOption) *Text {
	field := &Text{
		label: label,
	}

	for _, opt := range ops {
		opt(field)
	}

	return field
}

type Text struct {
	id           string
	label        string
	name         string
	value        string
	defaultValue string
	placeholder  string
	help         string
	validators   []func(string) error
	errors       []Error
	required     bool
}

func (t *Text) ID() string {
	return strings.ToLower(t.label)
}

func (t *Text) Name() string {
	return strings.ToLower(t.label)
}

func (t *Text) Label() template.HTML {
	str := `<label for="` + strings.ToLower(t.label) + `">`

	str += t.label
	if t.required {
		str += ` *`
	}

	str += `</label>`

	return template.HTML(str)
}

func (t *Text) Input() template.HTML {
	str := `<input type="text" id="` + strings.ToLower(t.label) + `"`

	str += ` name="` + strings.ToLower(t.label) + `"`
	str += ` value=""`

	if t.required {
		str += ` required`
	}

	str += `/>`

	return template.HTML(str)
}

// FIXME implement properly: wrapping div, errors ...
func (t *Text) Full() template.HTML {
	return t.Label() + t.Input()
}

func (t *Text) Value() string {
	return t.value
}

func (t *Text) SetValue(val string) {
	t.value = val
}

// Validate runs all validators on the field
func (t *Text) Validate() bool {
	for _, validator := range t.validators {
		if err := validator(t.value); err != nil {
			t.errors = append(t.errors, Error{Key: t.label, Message: err.Error()})
			return false
		}
	}

	return true
}

func (t *Text) Errors() []Error {
	return t.errors
}
