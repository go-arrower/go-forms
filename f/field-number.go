package f

import (
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
		hasValue: false,
	}

	for _, opt := range ops {
		opt.applyNumberOption(&field)
	}

	return field
}

type Number struct {
	base
	readonly bool
	hasValue bool
	value    int
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
	var value string

	if n.hasValue {
		value = strconv.Itoa(n.value)
	}

	str := `<input type="number" id="` + n.htmlID + `"`
	str += ` name="` + n.htmlName + `"`
	str += ` value="` + value + `"`

	if len(attr) > 0 && len(attr)%2 == 0 {
		if attr[0] == "class" {
			str += ` class="` + attr[1] + `"`
		}
	}

	if n.disabled {
		str += ` disabled`
	}

	if n.readonly {
		str += ` readonly`
	}

	str += `/>`

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

func (n *Number) Value() int {
	return n.value
}

func (n *Number) validate() bool {
	for _, validator := range n.validators {
		if err := validator(strconv.Itoa(n.value)); err != nil {
			n.errors = append(n.errors, Error{Key: n.label, Message: err.Error()})
			return false
		}
	}

	return true
}

func (n *Number) setValue(value any) {
	if value.(string) == "" {
		return
	}

	val, err := strconv.Atoi(value.(string))
	if err != nil {
		panic("go-forms: this field is implemented incorrectly: `Number` assumes int type for value: " + err.Error())
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
	//_ numberElement = (*placeholderOption)(nil)
	//_ numberElement = (*listOption)(nil)
	//_ numberElement = (*requiredValidator)(nil)
	//_ numberElement = (*autocompleteOption)(nil)
	//_ numberElement = (*formOption)(nil)
	//  max, min, step (float),
)
