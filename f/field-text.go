package f

import (
	"html/template"
	"strings"
)

// TextField constructs a new Text form inputElement.
func TextField(label string, ops ...textElement) Text {
	id := attrValue(label)

	field := Text{
		base: base{
			id:       id,
			label:    label,
			htmlName: id,
		},
	}

	for _, opt := range ops {
		opt.applyTextOption(&field)
	}

	return field
}

type Text struct {
	base
	maxlength   int
	minlength   int
	datalist    []string
	placeholder string
	help        string
	validators  []func(string) error
	errors      []Error
	required    bool
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
	hasList := len(t.datalist) > 0

	str := `<input type="text" id="` + t.id + `"`
	str += ` name="` + t.htmlName + `"`
	str += ` value="` + t.value + `"`

	if hasList {
		str += ` list="` + t.id + `-datalist"`
	}

	if t.required {
		str += ` required`
	}

	str += `/>`

	if hasList {
		str += `<datalist id="` + t.id + `-datalist">`

		for _, o := range t.datalist {
			str += `<option value="` + o + `"></option>`
		}

		str += `</datalist>`
	}

	return template.HTML(str)
}

// FIXME implement properly: wrapping div, errors ...
func (t *Text) Full() template.HTML {
	return t.Label() + t.Input()
}

func (t *Text) Value() string {
	return t.value
}

func (t *Text) setBase(base base) {
	t.base = base
}

// Validate runs all validators on the inputElement
func (t *Text) validate() bool {
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
