package f_test

import (
	"html/template"
	"reflect"
	"testing"

	"github.com/go-arrower/go-forms/f"
	"github.com/stretchr/testify/assert"
)

func TestSubmitButton(t *testing.T) {
	t.Parallel()

	t.Run("implements methods", func(t *testing.T) {
		t.Parallel()

		field := f.SubmitButton("")
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
	})

	t.Run("defaults", func(t *testing.T) {
		t.Parallel()

		html := template.HTML(`<input type="submit" value="Label"/>`)
		full := template.HTML(`<input type="submit" value="Label"/>`)

		field := f.SubmitButton("Label")

		assert.Equal(t, template.HTML(""), field.Label())
		assert.Equal(t, html, field.Input())
		assert.Equal(t, full, field.Full())
	})

	t.Run("with disabled", func(t *testing.T) {
		t.Parallel()

		field := f.SubmitButton("Label", f.WithDisabled())

		assert.Contains(t, field.Input(), " disabled")
		assert.Contains(t, field.Full(), " disabled")
	})
}
