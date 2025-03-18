package f

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"reflect"
	"strings"
)

type Field interface {
	Validate() bool

	// Name is the HTML attribute `name` used in the form.
	Name() string

	// SetValue takes the val from the http request via FormValue.
	// The field implementation can cast this to it's preferred type.
	// TODO check for complicated data like dropdown or select
	SetValue(val string) // could be any, as validation does cast already and does not work on string values alone
}

type FieldOption func(customField)

func WithID(id string) FieldOption {
	return func(field customField) {
		field.setID(id)
	}
}

func WithName(name string) FieldOption {
	return func(field customField) {
		field.setName(name)
	}
}

func WithValue(value any) FieldOption {
	return func(field customField) {
		field.setValue(value)
	}
}

func WithList(options []string) FieldOption {
	return func(field customField) {
		switch f := field.(type) {
		case *Text:
			f.datalist = options
		}
	}
}

func WithPlaceholder(placeholder string) FieldOption {
	return func(field customField) {
		switch f := field.(type) {
		case *Text:
			f.placeholder = placeholder
		default:
			slog.Log(context.Background(), slog.LevelDebug, "unsupported f Field type")
		}
	}
}

// customField needs to be implemented by every Field.
// This allows to easily add new fields. The only alternative is
// to switch on the type of field in wach FieldOption and set the values there.
// This would mean that for each new field, there would be multiple
// places that need to extend the code.
type customField interface {
	setID(string)
	setName(string)
	setValue(any)
}

// attrValue takes a label given by the user and converts in into a form,
// that is ready to use as a HTML id and name attribute.
// MDN recommends stricter rules, it is left to the developer.
// See: https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/id
func attrValue(label string) string {
	return strings.ReplaceAll(strings.ToLower(label), " ", "-")
}

// TODO consider: also return the errors, so the caller can make the decision to render the form or branch out to other logic
// OR a way to access the raw errors
func Validate(req *http.Request, form any) bool {
	if req == nil || form == nil {
		return false
	}

	formVal := reflect.ValueOf(form)
	for i := range formVal.NumField() {
		if err := req.ParseForm(); err != nil {
			return false
		}

		field, ok := formVal.Field(i).Interface().(Field)
		if !ok {
			return false
		}

		field.SetValue(req.FormValue(field.Name()))

		if !field.Validate() {
			return false
		}
	}

	return true
}

type Error struct {
	Key     string
	Message string
}

func (e Error) String() string {
	return e.Key + " " + e.Message
}

// --- --- ---
func Build() *Builder {
	return &Builder{Fields: map[string]Field{}}
}

type Builder struct {
	Fields map[string]Field
}

func (b *Builder) Text(label string, opts ...FieldOption) *Builder {
	b.Fields[label] = TextField(label, opts...)
	return b
}

func (b *Builder) Checkbox(label string, opts ...FieldOption) *Builder {
	b.Fields[label] = BooleanField(label, opts...)
	return b
}

func (b *Builder) Form() any {
	return b.Fields
}

func (b *Builder) Fields2() func(yield func(f Field) bool) {
	return func(yield func(f Field) bool) {
		for _, v := range b.Fields {
			if !yield(v) {
				return
			}
		}
	}
}

type Form map[string]Field

func New() *Form {
	return &Form{}
}

func (f *Form) Text(label string, opts ...FieldOption) *Form {
	// f[label] = TextField(label, opts...)
	(*f)[label] = TextField(label, opts...)
	return f
}

func (f *Form) Validate(req *http.Request) []Error {
	return nil
}

func Required() FieldOption {
	return func(f customField) {
		switch f := f.(type) {
		case *Text:
			f.required = true
			f.validators = append(f.validators, func(value string) error {
				if value == "" {
					return errors.New("is required")
				}

				return nil
			})
		}
	}
}
