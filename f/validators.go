package f

import "errors"

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

func (val requiredValidator) applyTextOption(f *Text) {
	f.required = true
	f.validators = append(f.validators, func(value string) error {
		if value == "" {
			return errors.New("is required")
		}

		return nil
	})
}
