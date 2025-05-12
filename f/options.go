package f

import (
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

func WithDisabled() disabledOption {
	return disabledOption(true)
}

func WithReadonly() readonlyOption {
	return readonlyOption(true)
}

func WithPlaceholder(placeholder string) placeholderOption {
	return placeholderOption(placeholder)
}

func WithList(options []string) listOption {
	return listOption(options)
}

func WithAutocomplete(autocomplete string) autocompleteOption {
	return autocompleteOption(autocomplete)
}

// WithSpellcheck hints the browser if spellchecking is desired.
// If not set, the behaviour is browser-defined.
func WithSpellcheck(enabled bool) spellcheckOption {
	if enabled {
		return "true"
	}

	return "false"
}

func WithAutocapitalize(capitalization autocapitalizeOption) autocapitalizeOption {
	return capitalization
}

func WithSize(size uint8) sizeOption {
	return sizeOption(size)
}

func WithTitle(title string) titleOption {
	return titleOption(title)
}

func WithForm(form string) formOption {
	return formOption(form)
}

func WithAutofocus(autofocus bool) autofocusOption {
	return autofocusOption(autofocus)
}

func WithMax(max time.Time) maxOption {
	return maxOption(max)
}

func WithStep(step float64) stepOption {
	return stepOption(step)
}

type (
	textElement interface {
		applyTextOption(f *Text)
	}

	numberElement interface {
		applyNumberOption(f *Number)
	}

	dateTimeLocalElement interface {
		applyDateTimeLocalOption(f *DateTimeLocal)
	}

	submitElement interface {
		applySubmitOption(f *Submit)
	}
)

type (
	idOption             string
	nameOption           string
	valueOption          struct{ value any }
	disabledOption       bool
	readonlyOption       bool
	placeholderOption    string
	listOption           []string
	autocompleteOption   string
	spellcheckOption     string
	autocapitalizeOption string
	sizeOption           uint8
	titleOption          string
	formOption           string
	tabindexOption       int //nolint:unused
	autofocusOption      bool
	maxOption            time.Time
	stepOption           float64
)

const (
	On         autocapitalizeOption = "on"
	Off        autocapitalizeOption = "off"
	None       autocapitalizeOption = "none"
	Sentences  autocapitalizeOption = "sentences"
	Words      autocapitalizeOption = "words"
	Characters autocapitalizeOption = "characters"
)

func (o idOption) applyTextOption(f *Text) {
	f.htmlID = string(o)
}

func (o idOption) applyNumberOption(f *Number) {
	f.htmlID = string(o)
}

func (o idOption) applyDateTimeLocalOption(f *DateTimeLocal) {
	f.htmlID = string(o)
}

func (o nameOption) applyTextOption(f *Text) {
	f.htmlName = string(o)
}

func (o nameOption) applyNumberOption(f *Number) {
	f.htmlName = string(o)
}

func (o nameOption) applyDateTimeLocalOption(f *DateTimeLocal) {
	f.htmlName = string(o)
}

func (o valueOption) applyTextOption(f *Text) {
	val, ok := o.value.(string)
	if !ok {
		panic("go-forms: WithValue for `Text` required a string")
	}

	f.value = val
}

func (o valueOption) applyNumberOption(f *Number) {
	val, ok := o.value.(float64)
	if !ok {
		iv, ok := o.value.(int)
		if !ok {
			panic("go-forms: WithValue for `Number` required an float64 or int")
		}
		val = float64(iv)
	}

	f.value = val
	f.hasValue = true
}
func (o valueOption) applyDateTimeLocalOption(f *DateTimeLocal) {
	val, ok := o.value.(time.Time)
	if !ok {
		panic("go-forms: WithValue for `DateTimeLocal` required a time.Time")
	}

	f.value = val.Format(browserLayout)
}

func (o disabledOption) applyTextOption(f *Text) {
	f.disabled = bool(o)
}

func (o disabledOption) applyNumberOption(f *Number) {
	f.disabled = bool(o)
}

func (o disabledOption) applyDateTimeLocalOption(f *DateTimeLocal) {
	f.disabled = bool(o)
}

func (o disabledOption) applySubmitOption(f *Submit) {
	f.disabled = bool(o)
}

func (o readonlyOption) applyTextOption(f *Text) {
	f.readonly = bool(o)
}

func (o readonlyOption) applyNumberOption(f *Number) {
	f.readonly = bool(o)
}

func (o readonlyOption) applyDateTimeLocalOption(f *DateTimeLocal) {
	f.readonly = bool(o)
}

func (o placeholderOption) applyTextOption(f *Text) {
	f.placeholder = string(o)
}

func (o placeholderOption) applyNumberOption(f *Number) {
	f.placeholder = string(o)
}

func (o listOption) applyTextOption(f *Text) {
	f.datalist = o
}

func (o listOption) applyNumberOption(f *Number) {
	f.datalist = o
}

func (o autocompleteOption) applyTextOption(f *Text) {
	f.autocomplete = string(o)
}

func (o autocompleteOption) applyNumberOption(f *Number) {
	f.autocomplete = string(o)
}

func (o spellcheckOption) applyTextOption(f *Text) {
	f.spellcheck = string(o)
}

func (o autocapitalizeOption) applyTextOption(f *Text) {
	f.autocapitalize = o
}

func (o sizeOption) applyTextOption(f *Text) {
	f.size = uint8(o)
}

func (o titleOption) applyTextOption(f *Text) {
	f.title = string(o)
}

func (o formOption) applyTextOption(f *Text) {
	f.form = string(o)
}

func (o formOption) applyNumberOption(f *Number) {
	f.form = string(o)
}

func (o autofocusOption) applyTextOption(f *Text) {
	f.autofocus = bool(o)
}

func (o maxOption) applyDateTimeLocalOption(f *DateTimeLocal) {
	// f.max = time.Time(o)
}

func (o stepOption) applyNumberOption(f *Number) {
	f.step = float64(o)
}
