package f_test

import (
	"html/template"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-arrower/go-forms/f"
)

func TestNumberField(t *testing.T) {
	t.Parallel()

	t.Run("implements methods", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("")
		tf := reflect.TypeOf(&field)

		method, exists := tf.MethodByName("Label")
		assert.True(t, exists, "*%s should implement method %s()", tf.Elem().Name(), "Label")
		assert.Equal(t, 1, method.Type.NumOut())
		assert.Equal(t, reflect.TypeOf(template.HTML("")), method.Type.Out(0))

		method, exists = tf.MethodByName("Input")
		assert.True(t, exists, "*%s should implement method %s()", tf.Elem().Name(), "Input")
		assert.Equal(t, 1, method.Type.NumOut())
		assert.Equal(t, reflect.TypeOf(template.HTML("")), method.Type.Out(0))

		method, exists = tf.MethodByName("Full")
		assert.True(t, exists, "*%s should implement method %s()", tf.Elem().Name(), "Full")
		assert.Equal(t, 1, method.Type.NumOut())
		assert.Equal(t, reflect.TypeOf(template.HTML("")), method.Type.Out(0))

		method, exists = tf.MethodByName("Errors")
		assert.True(t, exists, "*%s should implement method %s()", tf.Elem().Name(), "Errors")
		assert.Equal(t, 1, method.Type.NumOut())
		assert.Equal(t, reflect.TypeOf([]f.Error{}), method.Type.Out(0))

		method, exists = tf.MethodByName("Value")
		assert.True(t, exists, "*%s should implement method %s()", tf.Elem().Name(), "Value")
		assert.Equal(t, 1, method.Type.NumOut())
	})

	t.Run("defaults", func(t *testing.T) {
		t.Parallel()

		label := template.HTML(`<label for="label">Label</label>`)
		html := template.HTML(`<input type="number" id="label" name="label" value=""/>`)
		full := template.HTML(`<div><label for="label">Label</label><input type="number" id="label" name="label" value=""/></div>`)

		field := f.NumberField("Label")

		assert.Equal(t, label, field.Label())
		assert.Equal(t, html, field.Input())
		assert.Equal(t, full, field.Full())

		assert.Empty(t, field.Value())
		assert.Nil(t, field.Errors())
	})

	t.Run("with id", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label", f.WithID("my-id"))

		assert.Contains(t, field.Input(), `id="my-id"`)
		assert.Contains(t, field.Input(), `name="label"`, "option should only change id attribute")
		assert.Contains(t, field.Label(), `for="my-id"`, "option should only change id attribute")
	})

	t.Run("with name", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label", f.WithName("my-name"))

		assert.Contains(t, field.Input(), `name="my-name"`)
	})

	t.Run("with value", func(t *testing.T) {
		t.Parallel()

		t.Run("valid float", func(t *testing.T) {
			t.Parallel()

			field := f.NumberField("Label", f.WithValue(0.0))

			assert.Equal(t, 0.0, field.Value())
			assert.Contains(t, field.Input(), "0")
		})

		t.Run("valid int", func(t *testing.T) {
			t.Parallel()

			field := f.NumberField("Label", f.WithValue(0))

			assert.Equal(t, 0.0, field.Value())
			assert.Contains(t, field.Input(), "0")
		})

		t.Run("empty", func(t *testing.T) {
			t.Parallel()

			field := f.NumberField("Label", f.WithValue(0.0))
			assert.Equal(t, 0.0, field.Value())
		})

		t.Run("invalid", func(t *testing.T) {
			t.Parallel()

			assert.PanicsWithValue(t, "go-forms: WithValue for `Number` required an float64", func() {
				f.NumberField("Label", f.WithValue(true))
			})
		})
	})

	t.Run("with disabled", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label", f.WithDisabled())

		assert.Contains(t, field.Input(), " disabled")
		assert.Contains(t, field.Full(), " disabled")
	})

	t.Run("with readonly", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label", f.WithReadonly())

		assert.Contains(t, field.Input(), " readonly")
		assert.Contains(t, field.Full(), " readonly")
	})

	t.Run("with placeholder", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label", f.WithPlaceholder("my-placeholder"))

		assert.Contains(t, field.Input(), ` placeholder="my-placeholder"`)
		assert.Contains(t, field.Full(), ` placeholder="my-placeholder"`)
	})

	t.Run("with list", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label",
			f.WithID("my-id"),
			f.WithList([]string{"A", "B"}),
		)

		assert.Contains(t, field.Input(), ` list="my-id-datalist"`)
		assert.Contains(t, field.Input(), `<datalist id="my-id-datalist">`)
		assert.Contains(t, field.Input(), `<option value="A"></option>`)
		assert.Contains(t, field.Input(), `<option value="B"></option>`)
		assert.Contains(t, field.Input(), `</datalist>`)
	})

	t.Run("with autocomplete", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label", f.WithAutocomplete("on"))
		assert.Contains(t, field.Input(), `autocomplete="on"`)

		field = f.NumberField("Label", f.WithAutocomplete("off"))
		assert.Contains(t, field.Input(), `autocomplete="off"`)

		field = f.NumberField("Label", f.WithAutocomplete("family-name"))
		assert.Contains(t, field.Input(), `autocomplete="family-name"`)
	})

	t.Run("with form", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label", f.WithForm("my-form"))
		assert.Contains(t, field.Input(), ` form="my-form"`)
		assert.Contains(t, field.Full(), ` form="my-form"`)
	})

	t.Run("with input attributes", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label")

		assert.Contains(t, field.Input("class", "my-class"), ` class="my-class"`)
	})

	t.Run("validate", func(t *testing.T) {
		t.Parallel()

		form := f.New(struct{ Number f.Number }{})

		f.Validate(form, newRequest(""))
		assert.Equal(t, 0.0, form.Number.Value())
	})
}

func TestNumberField_Validation(t *testing.T) {
	t.Parallel()

	t.Run("html", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label", f.Required())
		assert.Contains(t, field.Label(), " *")
		assert.Contains(t, field.Input(), " required")
	})

	t.Run("inputs", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			postParams string
			pass       bool
			expValue   float64
		}{
			"empty":                     {"", false, 0},
			"input missing ":            {"1", false, 0},
			"input missing data":        {"f=", false, 0},
			"input has wrong data type": {"f=string-value", false, 0},
			"1":                         {"f=1", true, 1},
			"1.0":                       {"f=1.0", true, 1.0},
			"1.1":                       {"f=1.1", true, 1.1},
			"-0.1":                      {"f=-0.1", true, -0.1},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				form := f.New(struct{ F f.Number }{f.NumberField("")})

				assert.Equal(t, tt.pass, f.Validate(form, newRequest(tt.postParams)))
				assert.Equal(t, tt.expValue, form.F.Value())
			})
		}
	})

	t.Run("required", func(t *testing.T) {
		t.Parallel()

		tests := map[string]struct {
			postParams string
			pass       bool
			expValue   float64
		}{
			"empty":              {"", false, 0},
			"input missing ":     {"1", false, 0},
			"input missing data": {"f=", false, 0},
			"1":                  {"f=1", true, 1},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				form := f.New(struct{ F f.Number }{f.NumberField("", f.Required())})

				assert.Equal(t, tt.pass, f.Validate(form, newRequest(tt.postParams)))
				assert.Equal(t, tt.expValue, form.F.Value())

				if !tt.pass {
					assert.NotEmpty(t, form.F.Errors())
					assert.NotEmpty(t, form.F.Errors()[0].Key)
					assert.NotEmpty(t, form.F.Errors()[0].Message)
					t.Log(form.F.Errors())
				}
			})
		}
	})
}
