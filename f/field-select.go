package f

import (
	"html/template"
	"sort"
)

func SelectField(label string, choices any, ops ...any) Select {
	id := htmlAttr(label)

	if choices == nil {
		choices = []string{}
	}

	var options []choice

	switch choices := choices.(type) {
	case []string:
		for _, opt := range choices {
			options = append(options, choice{value: htmlAttr(opt), label: opt})
		}
	case [][]string:
		for _, ch := range choices {
			if len(ch) != 2 { //nolint:mnd // each option MUST have a value & label
				panic("invalid choices type, allowed is: [][2]string")
			}

			options = append(options, choice{value: ch[0], label: ch[1]})
		}
	case map[string]string:
		for value, label := range choices {
			options = append(options, choice{value: value, label: label})
		}

		sort.Slice(options, func(i, j int) bool {
			return options[i].label < options[j].label
		})
	default:
		panic("invalid choices type, allowed are: []string, [][]string, map[string][string], f.Optgroup")
	}

	field := Select{
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
		choices: options,
	}

	// for _, opt := range ops {
	// 	opt(field)
	// }

	return field
}

type choice struct {
	value string
	label string
}

type Select struct {
	base
	choices []choice
}

func (s *Select) Label() template.HTML {
	str := `<label for="` + htmlAttr(s.label) + `">`

	str += s.label
	if s.required {
		str += ` *`
	}

	str += `</label>`

	return template.HTML(str)
}

func (s *Select) Input() template.HTML {
	str := `<select id="` + s.id + `"`
	str += ` name="` + s.htmlName + `">`

	for _, option := range s.choices {
		str += `<option value="` + option.value + `">` + option.label + `</option>`
	}

	str += `</select>`

	return template.HTML(str)
}

func (s *Select) Full() template.HTML {
	str := `<div>` +
		s.Label() +
		s.Input()

	for _, e := range s.Errors() {
		str += `<span>` + template.HTML(e.Key) + ` ` + template.HTML(e.Message) + `</span>`
	}

	str += `</div>`

	return str
}

// TODO move base?
func (s *Select) Errors() []Error {
	return s.errors
}

func (s *Select) Value() string {
	return s.value
}

func (s *Select) validate() bool {
	for _, validator := range s.validators {
		if err := validator(s.value); err != nil {
			s.errors = append(s.errors, Error{Key: s.label, Message: err.Error()})
			return false
		}
	}

	return true
}

/*
values:
	optional value but mandatory name
	order
	<optgroup>
	hr
	legend
	attributes on the options like disabled or selected
*/

// withSelected("name")
// multiple
// size
// required, disabled, autofocus, autocomplete,form, name,
