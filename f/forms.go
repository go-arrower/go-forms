package f

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"reflect"
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

type FieldOption func(Field)

func WithID(id string) FieldOption {
	return func(field Field) {
		switch f := field.(type) {
		case *Text:
			f.id = id
		case *Boolean:
			f.id = id
		default:
			slog.Log(context.Background(), slog.LevelDebug, "unsupported f Field type")
		}
	}
}

func WithName(name string) FieldOption {
	return func(field Field) {
		switch f := field.(type) {
		case *Text:
			f.name = name
		case *Boolean:
			f.name = name
		default:
			slog.Log(context.Background(), slog.LevelDebug, "unsupported f Field type")
		}
	}
}

func WithPlaceholder(placeholder string) FieldOption {
	return func(field Field) {
		switch f := field.(type) {
		case *Text:
			f.placeholder = placeholder
		case *Boolean:
			f.placeholder = placeholder
		default:
			slog.Log(context.Background(), slog.LevelDebug, "unsupported f Field type")
		}
	}
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
	return func(f Field) {
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
