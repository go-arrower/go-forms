package f_test

import (
	"html/template"
	"reflect"
	"testing"

	"github.com/go-arrower/go-forms/f"
	"github.com/stretchr/testify/assert"
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

		t.Run("valid", func(t *testing.T) {
			t.Parallel()

			field := f.NumberField("Label", f.WithValue(0))

			assert.Equal(t, 0, field.Value())
			assert.Contains(t, field.Input(), "0")
		})

		t.Run("empty", func(t *testing.T) {
			t.Parallel()

			field := f.NumberField("Label", f.WithValue(0))
			assert.Equal(t, 0, field.Value())
		})

		t.Run("invalid", func(t *testing.T) {
			t.Parallel()

			assert.PanicsWithValue(t, "go-forms: WithValue for `Number` required an int", func() {
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

	t.Run("with input attributes", func(t *testing.T) {
		t.Parallel()

		field := f.NumberField("Label")

		assert.Contains(t, field.Input("class", "my-class"), ` class="my-class"`)
	})

	t.Run("validate", func(t *testing.T) {
		t.Parallel()

		form := f.New(struct{ Number f.Number }{})

		f.Validate(form, newRequest(""))
		assert.Equal(t, 0, form.Number.Value())
	})
}
