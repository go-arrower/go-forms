package f

import (
	"html/template"
	"strconv"
)

// TextField constructs a new Text form inputElement.
func TextField(label string, ops ...textElement) Text {
	id := htmlAttr(label)

	field := Text{
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
		maxlength:      0,
		minlength:      0,
		datalist:       nil,
		readonly:       false,
		placeholder:    "",
		autocomplete:   "",
		spellcheck:     "",
		autocapitalize: "",
		size:           0,
	}

	for _, opt := range ops {
		opt.applyTextOption(&field)
	}

	return field
}

type Text struct {
	base
	maxlength      int
	minlength      int
	datalist       []string
	readonly       bool
	placeholder    string
	autocomplete   string
	spellcheck     string
	autocapitalize autocapitalizeOption
	size           uint8
	// help        string
}

func (t *Text) Label() template.HTML {
	str := `<label for="` + htmlAttr(t.label) + `">`

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

	if t.required {
		str += ` required`
	}

	if t.disabled {
		str += ` disabled`
	}

	if t.readonly {
		str += ` readonly`
	}

	if t.placeholder != "" {
		str += ` placeholder="` + t.placeholder + `"`
	}

	if hasList {
		str += ` list="` + t.id + `-datalist"`
	}

	if t.autocomplete != "" {
		str += ` autocomplete="` + t.autocomplete + `"`
	}

	if t.spellcheck == "true" {
		str += ` spellcheck="true"`
	} else if t.spellcheck == "false" {
		str += ` spellcheck="false"`
	}

	switch t.autocapitalize {
	case On:
		str += ` autocapitalize="on"`
	case Off:
		str += ` autocapitalize="off"`
	case None:
		str += ` autocapitalize="none"`
	case Sentences:
		str += ` autocapitalize="sentences"`
	case Words:
		str += ` autocapitalize="words"`
	case Characters:
		str += ` autocapitalize="characters"`
	}

	if t.size > 0 {
		str += ` size="` + strconv.Itoa(int(t.size)) + `"`
	}

	if t.title != "" {
		str += ` title="` + t.title + `"`
	}

	if t.form != "" {
		str += ` form="` + t.form + `"`
	}

	if t.autofocus {
		str += ` autofocus`
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

func (t *Text) Full() template.HTML {
	str := `<div>` +
		t.Label() +
		t.Input()

	for _, e := range t.Errors() {
		str += `<span>` + template.HTML(e.Key) + ` ` + template.HTML(e.Message) + `</span>`
	}

	str += `</div>`

	return str
}

func (t *Text) Errors() []Error {
	return t.errors
}

func (t *Text) Value() string {
	return t.value
}

func (t *Text) validate() bool {
	for _, validator := range t.validators {
		if err := validator(t.value); err != nil {
			t.errors = append(t.errors, Error{Key: t.label, Message: err.Error()})
			return false
		}
	}

	return true
}

var (
	_ textElement = (*idOption)(nil)
	_ textElement = (*nameOption)(nil)
	_ textElement = (*valueOption)(nil)
	_ textElement = (*disabledOption)(nil)
	_ textElement = (*readonlyOption)(nil)
	_ textElement = (*placeholderOption)(nil)
	_ textElement = (*listOption)(nil)
	_ textElement = (*requiredValidator)(nil)
	// _ textElement = (*patternValidator)(nil)
	_ textElement = (*maxlengthValidator)(nil)
	_ textElement = (*minlengthValidator)(nil)
	_ textElement = (*autocompleteOption)(nil)
	_ textElement = (*spellcheckOption)(nil)
	_ textElement = (*autocapitalizeOption)(nil)
	_ textElement = (*sizeOption)(nil)
	_ textElement = (*titleOption)(nil)
	_ textElement = (*formOption)(nil)
	// _ textElement = (*tabindexOption)(nil)
	_ textElement = (*autofocusOption)(nil)
)
