package f_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-arrower/go-forms/f"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("initialise defaults", func(t *testing.T) {
		t.Parallel()

		form := f.New(myForm{})

		assert.Contains(t, form.Firstname.Label(), "Firstname")
		assert.Contains(t, form.Lastname.Label(), "Lastname")
		// TODO assert on other defaults like ID, name etc
	})

	t.Run("overwrites", func(t *testing.T) {
		t.Parallel()

		form := f.New(myForm{
			Lastname: f.TextField("Last Name"),
		})

		assert.Contains(t, form.Firstname.Label(), "Firstname")
		assert.Contains(t, form.Lastname.Label(), "Last Name")
	})

	t.Run("invalid input", func(t *testing.T) {
		t.Parallel()

		textField := f.TextField("")

		tests := []struct {
			form     any
			expected string
		}{
			{"", "string"},
			{0, "int"},
			{0.1, "float64"},
			{nil, "nil"},
			{fmt.Stringer(nil), "nil"},
			{textField, "Text"},
			{&textField, ""},
		}

		for _, tt := range tests {
			t.Run(tt.expected, func(t *testing.T) {
				t.Parallel()

				assert.PanicsWithValue(t,
					fmt.Sprintf("could not create new go-form: %s is not a valid form", tt.expected),
					func() {
						f.New(tt.form)
					},
				)
			})
		}
	})

	t.Run("handle unknown input fields", func(t *testing.T) {
		t.Parallel()

		type myForm struct {
			Firstname f.Text
			Text      string
		}

		assert.PanicsWithValue(t,
			"field Text of myForm is not an inputElement",
			func() {
				f.New(myForm{})
			})
	})
}

func TestValidate(t *testing.T) {
	t.Parallel()

	t.Run("validation fails", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			f        any
			r        *http.Request
			expValid bool
		}{
			"nil request":      {struct{}{}, nil, false},
			"nil form":         {nil, newRequest(""), false},
			"form not pointer": {myForm{}, newRequest(""), false},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				assert.Equal(t, tt.expValid, f.Validate(tt.f, tt.r))
			})
		}
	})

	t.Run("validation panics", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			f          any
			r          *http.Request
			expMessage string
		}{
			"invalid form": {&invalidForm{}, newRequest(""), "field Name of invalidForm is not an inputElement"},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				assert.PanicsWithValue(t, tt.expMessage, func() {
					f.Validate(tt.f, tt.r)
				})
			})
		}
	})

	t.Run("populated form", func(t *testing.T) {
		t.Parallel()

		form := f.New(myForm{Firstname: f.TextField("param")})

		assert.True(t, f.Validate(form, newRequest("param=value&lastname=lastname")))
		assert.Contains(t, form.Firstname.Label(), "param")
		assert.Equal(t, "value", form.Firstname.Value())

		assert.Contains(t, form.Lastname.Label(), "Lastname")
		assert.Equal(t, "lastname", form.Lastname.Value())
	})
}

func newRequest(data string) *http.Request {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req
}

type myForm struct {
	Firstname f.Text
	Lastname  f.Text
}

type invalidForm struct {
	Name string
}
