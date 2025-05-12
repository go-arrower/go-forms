package f

import (
	"errors"
	"html/template"
	"strconv"
)

func NumberField(label string, ops ...numberElement) Number {
	id := htmlAttr(label)

	field := Number{
		base: base{
			htmlID:       id,
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
		readonly:     false,
		hasValue:     false,
		value:        0,
		placeholder:  "",
		datalist:     nil,
		autocomplete: "",
	}

	for _, opt := range ops {
		opt.applyNumberOption(&field)
	}

	return field
}

type Number struct {
	base
	readonly bool
	// hasValue is used to distinguish between the default value 0 for int and empty
	// this is necessary, to differentiate when checking for the required option.
	hasValue     bool
	value        float64
	placeholder  string
	datalist     []string
	autocomplete string
	step         float64
}

func (n *Number) Label() template.HTML {
	str := `<label for="` + n.htmlID + `">`

	str += n.label
	if n.required {
		str += ` *`
	}

	str += `</label>`

	return template.HTML(str)
}

func (n *Number) Input(attr ...string) template.HTML {
	hasList := len(n.datalist) > 0

	var value string

	if n.hasValue {
		value = strconv.FormatFloat(n.value, 'f', -1, 64)
	}

	str := `<input type="number" id="` + n.htmlID + `"`
	str += ` name="` + n.htmlName + `"`
	str += ` value="` + value + `"`

	if n.step != 0 {
		str += ` step="` + strconv.FormatFloat(n.step, 'f', -1, 64) + `"`
	}

	if len(attr) > 0 && len(attr)%2 == 0 {
		if attr[0] == "class" {
			str += ` class="` + attr[1] + `"`
		}
	}

	if n.required {
		str += ` required`
	}

	if n.disabled {
		str += ` disabled`
	}

	if n.readonly {
		str += ` readonly`
	}

	if n.placeholder != "" {
		str += ` placeholder="` + n.placeholder + `"`
	}

	if n.autocomplete != "" {
		str += ` autocomplete="` + n.autocomplete + `"`
	}

	if n.form != "" {
		str += ` form="` + n.form + `"`
	}

	if hasList {
		str += ` list="` + n.htmlID + `-datalist"`
	}

	str += `/>`

	if hasList {
		str += `<datalist id="` + n.htmlID + `-datalist">`

		for _, o := range n.datalist {
			str += `<option value="` + o + `"></option>`
		}

		str += `</datalist>`
	}

	return template.HTML(str)
}

func (n *Number) Full() template.HTML {
	str := `<div>` +
		n.Label() +
		n.Input()

	for _, e := range n.Errors() {
		str += `<span>` + template.HTML(e.Key) + ` ` + template.HTML(e.Message) + `</span>`
	}

	str += `</div>`

	return str
}
func (n *Number) Errors() []Error {
	return n.errors
}

func (n *Number) Value() float64 {
	return n.value
}

func (n *Number) validate() bool {
	for _, validator := range n.validators {
		if err := validator(n, strconv.FormatFloat(n.value, 'f', 10, 64)); err != nil {
			n.errors = append(n.errors, Error{Key: n.label, Message: err.Error()})
			return false
		}
	}

	return true
}

// todo does the http request always give a string? than make this clear here, so the method does only parsing not casting
func (n *Number) setValue(value string) {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		n.validators = append(n.validators, func(field any, value string) error {
			return errors.New("value must be float64")
		})
	}

	n.value = val
	n.hasValue = true
}

var (
	_ numberElement = (*idOption)(nil)
	_ numberElement = (*nameOption)(nil)
	_ numberElement = (*valueOption)(nil)
	_ numberElement = (*disabledOption)(nil)
	_ numberElement = (*readonlyOption)(nil)
	_ numberElement = (*placeholderOption)(nil)
	_ numberElement = (*listOption)(nil)
	_ numberElement = (*requiredValidator)(nil)
	_ numberElement = (*autocompleteOption)(nil)
	_ numberElement = (*formOption)(nil)
	_ numberElement = (*stepOption)(nil)
	//  max, min, step (float),
)
