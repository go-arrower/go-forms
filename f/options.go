package f

import (
	"context"
	"log/slog"
)

type FieldOption func(inputElement)

func WithID(id string) FieldOption {
	return func(field inputElement) {
		field.setID(id)
	}
}

func WithName(name string) FieldOption {
	return func(field inputElement) {
		field.setName(name)
	}
}

func WithValue(value any) FieldOption {
	return func(field inputElement) {
		field.setValue(value)
	}
}

func WithList(options []string) FieldOption {
	return func(field inputElement) {
		switch f := field.(type) {
		case *Text:
			f.datalist = options
		}
	}
}

func WithPlaceholder(placeholder string) FieldOption {
	return func(field inputElement) {
		switch f := field.(type) {
		case *Text:
			f.placeholder = placeholder
		default:
			slog.Log(context.Background(), slog.LevelDebug, "unsupported f inputElement type")
		}
	}
}
