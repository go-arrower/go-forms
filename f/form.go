package f

import (
	"net/http"
	"reflect"
	"strings"
)

// New takes a form struct and initialises it with the default values, so it
// is ready to use in a html template.
// form has to be a struct where all fields are supported fields of go-form.
// In case of an error, it panics.
func New[F any](form F) *F {
	ft := reflect.TypeOf(form)
	fv := reflect.ValueOf(&form).Elem()

	if ft == nil {
		panic("could not create new go-form: nil is not a valid form")
	}

	if fv.Kind() != reflect.Struct {
		panic("could not create new go-form: " + reflect.TypeOf(form).Name() + " is not a valid form")
	}

	for i := range fv.NumField() {
		if fv.Field(i).IsZero() {
			field, ok := fv.Field(i).Addr().Interface().(inputElement)
			if !ok {
				panic("field " + ft.Field(i).Name + " of " + reflect.TypeOf(form).Name() + " is not an inputElement")
			}

			id := htmlAttr(ft.Field(i).Name)
			field.setBase(base{
				id:           id,
				label:        ft.Field(i).Name,
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
			})
		}
	}

	return &form
}

// Validate validates a given form against a http request.
// The data from the request is set to the field of the form,
// so if validation passes, the form is ready to use.
func Validate(form any, req *http.Request) bool {
	if req == nil || form == nil {
		return false
	}

	fv := reflect.ValueOf(form)
	if fv.Kind() != reflect.Ptr {
		return false
	}

	if err := req.ParseForm(); err != nil {
		return false
	}

	fv = fv.Elem()
	for i := range fv.NumField() {
		field, ok := fv.Field(i).Addr().Interface().(inputElement)
		if !ok {
			ft := reflect.TypeOf(form).Elem()
			panic("field " + ft.Field(i).Name + " of " + ft.Name() + " is not an inputElement")
		}

		field.setValue(req.FormValue(field.name()))

		if !field.validate() {
			return false
		}
	}

	return true
}

type inputElement interface {
	validate() bool

	setBase(base base)

	setID(id string)

	// name is the HTML attribute `name` used in the form.
	name() string
	setName(name string)

	// setValue takes the val from the http request via FormValue.
	// The inputElement implementation can cast this to it's preferred type.
	setValue(value any)
}

// htmlAttr takes a label given by the user and converts in into a form,
// that is ready to use as a HTML id and name attribute.
// MDN recommends stricter rules, it is left to the developer.
// See: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/id
func htmlAttr(label string) string {
	return strings.ReplaceAll(strings.ToLower(label), " ", "-")
}
