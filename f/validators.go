package f

import (
	"errors"
	"regexp"
)

var errIsRequired = errors.New("is required")

// Error is a single validation error.
type Error struct {
	Key     string
	Message string
}

func (e Error) String() string {
	return e.Key + " " + e.Message
}

func Required() requiredValidator {
	return requiredValidator(true)
}

// WithPattern parses the expression
// It ignores invalid patterns silently, as specified by MDN:
// "If the pattern attribute is present but is not specified or is invalid,
// no regular expression is applied and this attribute is ignored completely."
// see:https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#pattern
func WithPattern(pattern string) patternValidator {
	panic("implement me")
}

func WithMaxLength(max int) maxlengthValidator {
	return maxlengthValidator(max)
}

func WithMinLength(min int) minlengthValidator {
	return minlengthValidator(min)
}

type (
	requiredValidator  bool
	patternValidator   regexp.Regexp
	maxlengthValidator int
	minlengthValidator int
)

func (val requiredValidator) applyTextOption(f *Text) {
	f.required = true
	f.validators = append(f.validators, func(value string) error {
		if value == "" {
			return errIsRequired
		}

		return nil
	})
}

func (o maxlengthValidator) applyTextOption(f *Text) {
	f.maxlength = int(o)
}

func (o minlengthValidator) applyTextOption(f *Text) {
	f.minlength = int(o)
}
