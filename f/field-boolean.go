package f

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

func BooleanField(label string, ops ...any) *Boolean {
	return &Boolean{label: label}
}

type Boolean struct {
	id           string
	label        string
	name         string
	value        bool
	defaultValue string
	placeholder  string
	validators   []func(bool) error
	errors       []Error
}

func (t *Boolean) ID() string {
	return strings.ToLower(t.label)
}

func (t *Boolean) Name() string {
	return strings.ToLower(t.label)
}

func (t *Boolean) Label() template.HTML {
	return template.HTML(
		fmt.Sprintf(`<label for="%s">%s</label>`,
			strings.ToLower(t.label), t.label))
}

func (t *Boolean) Input() template.HTML {
	return template.HTML(
		fmt.Sprintf(`<input type="text" id="%s" name="%s" value=""/>`,
			strings.ToLower(t.label), strings.ToLower(t.label)))
}

func (t *Boolean) Value() bool {
	return t.value
}

func (t *Boolean) SetValue(val string) {
	t.value, _ = strconv.ParseBool(val) // TODO test if http forms use the values of ParseBool
}

func (t *Boolean) Validate() bool {
	for _, validator := range t.validators {
		if err := validator(t.value); err != nil {
			t.errors = append(t.errors, Error{Message: err.Error()})
			return false
		}
	}
	return true
}
