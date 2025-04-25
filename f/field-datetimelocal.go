package f

import (
	"html/template"
	"time"
)

func DateTimeLocalField(label string, ops ...dateTimeLocalElement) DateTimeLocal {
	id := htmlAttr(label)

	field := DateTimeLocal{
		base: base{
			id:           id,
			label:        label,
			htmlName:     id,
			value:        "",
			validators:   nil,
			errors:       nil,
			required:     false,
			disabled:     false,
			defaultValue: "",
			title:        "",
			form:         "",
			autofocus:    false,
		},
		readonly: false,
	}

	for _, opt := range ops {
		opt.applyDateTimeLocalOption(&field)
	}

	return field
}

type DateTimeLocal struct {
	base
	readonly bool
}

func (t *DateTimeLocal) Label() template.HTML {
	str := `<label for="` + htmlAttr(t.label) + `">`
	str += t.label
	str += `</label>`

	return template.HTML(str)
}

func (t *DateTimeLocal) Input() template.HTML {
	str := `<input type="datetime-local" id="` + t.id + `"`
	str += ` name="` + t.htmlName + `"`
	str += ` value="` + t.value + `"`

	if t.disabled {
		str += ` disabled`
	}

	if t.readonly {
		str += ` readonly`
	}

	str += `/>`

	return template.HTML(str)
}

func (t *DateTimeLocal) Full() template.HTML {
	str := `<div>` +
		t.Label() +
		t.Input()

	for _, e := range t.Errors() {
		str += `<span>` + template.HTML(e.Key) + ` ` + template.HTML(e.Message) + `</span>`
	}

	str += `</div>`

	return str
}

func (t *DateTimeLocal) Errors() []Error {
	return t.errors
}

func (t *DateTimeLocal) Value() time.Time {
	time, _ := time.Parse(time.RFC3339Nano, t.value)
	return time
}

// TODO can this be moved to base?
func (t *DateTimeLocal) validate() bool {
	for _, validator := range t.validators {
		if err := validator(t.value); err != nil {
			t.errors = append(t.errors, Error{Key: t.label, Message: err.Error()})
			return false
		}
	}

	return true
}

var (
	_ dateTimeLocalElement = (*idOption)(nil)
	_ dateTimeLocalElement = (*nameOption)(nil)
	_ dateTimeLocalElement = (*valueOption)(nil)
	_ dateTimeLocalElement = (*disabledOption)(nil)
	_ dateTimeLocalElement = (*readonlyOption)(nil)
	// _ dateTimeLocalElement = (*listOption)(nil)
	// _ dateTimeLocalElement = (*requiredValidator)(nil)
	// _ dateTimeLocalElement = (*autocompleteOption)(nil)
	// _ dateTimeLocalElement = (*autocapitalizeOption)(nil)
	// _ dateTimeLocalElement = (*titleOption)(nil)
	// _ dateTimeLocalElement = (*formOption)(nil)
	// _ dateTimeLocalElement = (*maxOption)(nil)
	// _ dateTimeLocalElement = (*minOption)(nil)
	// _ dateTimeLocalElement = (*stepOption)(nil)
	// _ dateTimeLocalElement = (*tabindexOption)(nil)
)
