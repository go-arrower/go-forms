package f_test

import (
	"html/template"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-arrower/go-forms/f"
)

func TestDateTimeLocalField(t *testing.T) {
	t.Parallel()

	t.Run("implements methods", func(t *testing.T) {
		t.Parallel()

		field := f.DateTimeLocalField("")
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
		html := template.HTML(`<input type="datetime-local" id="label" name="label" value=""/>`)
		full := template.HTML(`<div><label for="label">Label</label><input type="datetime-local" id="label" name="label" value=""/></div>`)

		field := f.DateTimeLocalField("Label")

		assert.Equal(t, label, field.Label())
		assert.Equal(t, html, field.Input())
		assert.Equal(t, full, field.Full())

		assert.True(t, field.Value().IsZero())
		assert.Nil(t, field.Errors())
	})

	t.Run("with id", func(t *testing.T) {
		t.Parallel()

		field := f.DateTimeLocalField("Label", f.WithID("my-id"))

		assert.Contains(t, field.Input(), `id="my-id"`)
		assert.Contains(t, field.Input(), `name="label"`, "option should only change id attribute")
		assert.Contains(t, field.Label(), `for="label"`, "option should only change id attribute")
	})

	t.Run("with name", func(t *testing.T) {
		t.Parallel()

		field := f.DateTimeLocalField("Label", f.WithName("my-name"))

		assert.Contains(t, field.Input(), `name="my-name"`)
	})

	t.Run("with value", func(t *testing.T) {
		t.Parallel()

		t.Run("valid", func(t *testing.T) {
			t.Parallel()

			now := time.Now()
			field := f.DateTimeLocalField("Label", f.WithValue(now))

			assert.True(t, now.Equal(field.Value()))
			assert.Contains(t, field.Input(), now.Format(time.RFC3339Nano))
		})

		t.Run("empty", func(t *testing.T) {
			t.Parallel()

			field := f.DateTimeLocalField("Label", f.WithValue(time.Time{}))
			assert.Equal(t, time.Time{}, field.Value())
		})

		t.Run("invalid", func(t *testing.T) {
			t.Parallel()

			assert.PanicsWithValue(t, "go-forms: WithValue for `DateTimeLocal` required a time.Time", func() {
				f.DateTimeLocalField("Label", f.WithValue(true))
			})
		})
	})

	t.Run("with disabled", func(t *testing.T) {
		t.Parallel()

		field := f.DateTimeLocalField("Label", f.WithDisabled())

		assert.Contains(t, field.Input(), " disabled")
		assert.Contains(t, field.Full(), " disabled")
	})

	t.Run("with readonly", func(t *testing.T) {
		t.Parallel()

		field := f.DateTimeLocalField("Label", f.WithReadonly())

		assert.Contains(t, field.Input(), " readonly")
		assert.Contains(t, field.Full(), " readonly")
	})
}
