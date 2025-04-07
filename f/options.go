package f

import (
	"regexp"
	"time"
)

func WithID(id string) idOption {
	return idOption(id)
}

func WithName(name string) nameOption {
	return nameOption(name)
}

func WithValue(value any) valueOption {
	return valueOption{value: value}
}

func WithPlaceholder(placeholder string) placeholderOption {
	return placeholderOption(placeholder)
}

func WithList(options []string) listOption {
	return listOption(options)
}

func WithMaxLength(max int) maxlengthValidator {
	return maxlengthValidator(max)
}

func WithMinLength(min int) minlengthValidator {
	return minlengthValidator(min)
}

func WithMax(max time.Time) maxOption {
	return maxOption(max)
}

type (
	textElement interface {
		applyTextOption(f *Text)
	}

	dateTimeLocalElement interface {
		applyDateTimeLocalOption(f *DateTimeLocal)
	}
)

var (
	_ textElement = (*idOption)(nil)
	_ textElement = (*nameOption)(nil)
	_ textElement = (*valueOption)(nil)
	// _ textElement = (*disabledOption)(nil)
	// _ textElement = (*readonlyOption)(nil)
	_ textElement = (*placeholderOption)(nil)
	_ textElement = (*listOption)(nil)
	_ textElement = (*requiredValidator)(nil)
	_ textElement = (*maxlengthValidator)(nil)
	_ textElement = (*minlengthValidator)(nil)
	// _ textElement = (*patternValidator)(nil)
	// _ textElement = (*autocompleteOption)(nil)
	// _ textElement = (*spellcheckOption)(nil)
	// _ textElement = (*autocapitalizeOption)(nil)
	// _ textElement = (*sizeOption)(nil)
	// _ textElement = (*titleOption)(nil)
	// _ textElement = (*formOption)(nil)
	// _ textElement = (*tabindexOption)(nil)
	// _ textElement = (*autofocusOption)(nil)
)

type (
	idOption             string
	nameOption           string
	valueOption          struct{ value any }
	disabledOption       bool
	readonlyOption       bool
	placeholderOption    string
	listOption           []string
	requiredValidator    bool
	maxlengthValidator   int
	minlengthValidator   int
	patternValidator     regexp.Regexp
	autocompleteOption   bool
	spellcheckOption     bool
	autocapitalizeOption bool
	sizeOption           int
	titleOption          string
	formOption           string
	tabindexOption       int
	autofocusOption      bool
	maxOption            time.Time
)

func (o idOption) applyTextOption(f *Text) {
	f.id = string(o)
}

func (o idOption) applyDateTimeLocalOption(f *DateTimeLocal) {
	f.id = string(o)
}

func (o nameOption) applyTextOption(f *Text) {
	f.htmlName = string(o)
}

func (o valueOption) applyTextOption(f *Text) {
	f.value = o.value.(string)
}

func (o placeholderOption) applyTextOption(f *Text) {
	f.placeholder = string(o)
}

func (o listOption) applyTextOption(f *Text) {
	f.datalist = o
}

func (o maxlengthValidator) applyTextOption(f *Text) {
	f.maxlength = int(o)
}

func (o minlengthValidator) applyTextOption(f *Text) {
	f.minlength = int(o)
}

func (o maxOption) applyDateTimeLocalOption(f *DateTimeLocal) {
	// f.max = time.Time(o)
}
