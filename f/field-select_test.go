package f_test

import (
	"html/template"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-arrower/go-forms/f"
)

func TestSelectField(t *testing.T) {
	t.Parallel()

	t.Run("implements methods", func(t *testing.T) {
		t.Parallel()

		field := f.SelectField("", nil)
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
		html := template.HTML(`<select id="label" name="label"><option value="o-1">O 1</option><option value="o2">O2</option></select>`)
		full := template.HTML(`<div><label for="label">Label</label><select id="label" name="label"><option value="o-1">O 1</option><option value="o2">O2</option></select></div>`)

		field := f.SelectField("Label", []string{"O 1", "O2"})

		assert.Equal(t, label, field.Label())
		assert.Equal(t, html, field.Input())
		assert.Equal(t, full, field.Full())

		assert.Empty(t, field.Value())
		assert.Nil(t, field.Errors())
	})

	t.Run("input formats", func(t *testing.T) {
		t.Parallel()

		field := f.SelectField("", []string{"O 1", "O2"})
		assert.Contains(t, field.Input(), `<option value="o-1">O 1</option>`)
		assert.Contains(t, field.Input(), `<option value="o2">O2</option>`)

		field = f.SelectField("", [][]string{{"value", "label"}, {"o-1", "O 1"}})
		assert.Contains(t, field.Input(), `<option value="value">label</option>`)
		assert.Contains(t, field.Input(), `<option value="o-1">O 1</option>`)

		assert.PanicsWithValue(t, "invalid choices type, allowed is: [][2]string", func() {
			field = f.SelectField("", [][]string{{"value", "label", "other"}})
		})

		field = f.SelectField("", map[string]string{
			"value": "label",
			"o-1":   "O 1",
		})
		assert.Contains(t, field.Input(), `<option value="value">label</option>`)
		assert.Contains(t, field.Input(), `<option value="o-1">O 1</option>`)
		assert.Contains(t, field.Input(),
			`<option value="o-1">O 1</option><option value="value">label</option>`,
			"map has random order but UI should be deterministic; sort by label",
		)

		// field = f.SelectField("", // TODO
		// 	Optgroup(
		// 		opt("", ""),
		// 		G("g0",
		// legend(), // replaces optgroup label
		// opt("", "" /*selected(), disabled()*/),
		// opt("", ""),
		// ),
		// Hr(),
		// G("g1", opt("", "", nil))),
		// f.WithID("my-id"),
		// )

		field = f.SelectField("", nil)
		assert.NotContains(t, field.Input(), `<option`)

		field = f.SelectField("", []string{})
		assert.NotContains(t, field.Input(), `<option`)

		field = f.SelectField("", [][]string{})
		assert.NotContains(t, field.Input(), `<option`)

		field = f.SelectField("", map[string]string{})
		assert.NotContains(t, field.Input(), `<option`)

		assert.PanicsWithValue(t, "invalid choices type, allowed are: []string, [][]string, map[string][string], f.Optgroup", func() {
			field = f.SelectField("", time.Time{})
		})
	})
}

type optType struct{}

func opt(value, label string, qualifiers ...any) optType {
	return optType{}
}

func Optgroup(g ...Group) any {
	return nil
}

func G(name string, o ...optType) any {
	return struct{ n string }{"NAME " + name}
}

type (
	Group any
)
