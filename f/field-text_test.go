package f_test

import (
	"bytes"
	"html/template"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-arrower/go-forms/f"
)

func TestTextField(t *testing.T) {
	t.Parallel()

	t.Run("implements methods", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("")
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
		html := template.HTML(`<input type="text" id="label" name="label" value=""/>`)
		full := template.HTML(`<div><label for="label">Label</label><input type="text" id="label" name="label" value=""/></div>`)

		field := f.TextField("Label")

		assert.Equal(t, label, field.Label())
		assert.Equal(t, html, field.Input())
		assert.Equal(t, full, field.Full())

		assert.Empty(t, field.Value())
		assert.Nil(t, field.Errors())
	})

	t.Run("with id", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithID("my-id"))

		assert.Contains(t, field.Input(), `id="my-id"`)
		assert.Contains(t, field.Input(), `name="label"`, "option should only change id attribute")
		assert.Contains(t, field.Label(), `for="my-id"`, "option should only change id attribute")
	})

	t.Run("with name", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithName("my-name"))

		assert.Contains(t, field.Input(), `name="my-name"`)
	})

	t.Run("with value", func(t *testing.T) {
		t.Parallel()

		t.Run("valid", func(t *testing.T) {
			t.Parallel()

			field := f.TextField("Label", f.WithValue("my-value"))

			assert.Equal(t, "my-value", field.Value())
			assert.Contains(t, field.Input(), "my-value")
		})

		t.Run("empty", func(t *testing.T) {
			t.Parallel()

			field := f.TextField("Label", f.WithValue(""))
			assert.Equal(t, "", field.Value())
		})

		t.Run("invalid", func(t *testing.T) {
			t.Parallel()

			assert.PanicsWithValue(t, "go-forms: WithValue for `Text` required a string", func() {
				f.TextField("Label", f.WithValue(true))
			})
		})
	})

	t.Run("with disabled", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithDisabled())

		assert.Contains(t, field.Input(), " disabled")
		assert.Contains(t, field.Full(), " disabled")
	})

	t.Run("with readonly", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithReadonly())

		assert.Contains(t, field.Input(), " readonly")
		assert.Contains(t, field.Full(), " readonly")
	})

	t.Run("with placeholder", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithPlaceholder("my-placeholder"))

		assert.Contains(t, field.Input(), ` placeholder="my-placeholder"`)
		assert.Contains(t, field.Full(), ` placeholder="my-placeholder"`)
	})

	t.Run("with list", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label",
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

		field := f.TextField("Label", f.WithAutocomplete("on"))
		assert.Contains(t, field.Input(), `autocomplete="on"`)

		field = f.TextField("Label", f.WithAutocomplete("off"))
		assert.Contains(t, field.Input(), `autocomplete="off"`)

		field = f.TextField("Label", f.WithAutocomplete("family-name"))
		assert.Contains(t, field.Input(), `autocomplete="family-name"`)
	})

	t.Run("with spellcheck", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithSpellcheck(true))

		assert.Contains(t, field.Input(), ` spellcheck="true"`)
		assert.Contains(t, field.Full(), ` spellcheck="true"`)

		field = f.TextField("Label", f.WithSpellcheck(false))

		assert.Contains(t, field.Input(), ` spellcheck="false"`)
		assert.Contains(t, field.Full(), ` spellcheck="false"`)

		field = f.TextField("Label")

		assert.NotContains(t, field.Input(), `spellcheck`)
		assert.NotContains(t, field.Full(), `spellcheck`)
	})

	t.Run("with autocapitalize", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithAutocapitalize(f.On))
		assert.Contains(t, field.Input(), ` autocapitalize="on"`)
		assert.Contains(t, field.Full(), ` autocapitalize="on"`)

		field = f.TextField("Label", f.WithAutocapitalize(f.Off))
		assert.Contains(t, field.Input(), ` autocapitalize="off"`)
		assert.Contains(t, field.Full(), ` autocapitalize="off"`)

		field = f.TextField("Label", f.WithAutocapitalize(f.None))
		assert.Contains(t, field.Input(), ` autocapitalize="none"`)
		assert.Contains(t, field.Full(), ` autocapitalize="none"`)

		field = f.TextField("Label", f.WithAutocapitalize(f.Sentences))
		assert.Contains(t, field.Input(), ` autocapitalize="sentences"`)
		assert.Contains(t, field.Full(), ` autocapitalize="sentences"`)

		field = f.TextField("Label", f.WithAutocapitalize(f.Words))
		assert.Contains(t, field.Input(), ` autocapitalize="words"`)
		assert.Contains(t, field.Full(), ` autocapitalize="words"`)

		field = f.TextField("Label", f.WithAutocapitalize(f.Characters))
		assert.Contains(t, field.Input(), ` autocapitalize="characters"`)
		assert.Contains(t, field.Full(), ` autocapitalize="characters"`)

		field = f.TextField("Label", f.WithAutocapitalize("my-custom-value"))
		assert.NotContains(t, field.Input(), `autocapitalize`)
		assert.NotContains(t, field.Full(), `autocapitalize`)
	})

	t.Run("with size", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithSize(10))
		assert.Contains(t, field.Input(), ` size="10"`)
		assert.Contains(t, field.Full(), ` size="10"`)
	})

	t.Run("with title", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithTitle("my-title"))
		assert.Contains(t, field.Input(), ` title="my-title"`)
		assert.Contains(t, field.Full(), ` title="my-title"`)
	})

	t.Run("with form", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithForm("my-form"))
		assert.Contains(t, field.Input(), ` form="my-form"`)
		assert.Contains(t, field.Full(), ` form="my-form"`)
	})

	t.Run("with form", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithAutofocus(true))
		assert.Contains(t, field.Input(), ` autofocus`)
		assert.Contains(t, field.Full(), ` autofocus`)
	})
}

// This test case is to ensure the inputElement works in its entirety.
// Rendering it through the arrower renderer ensures behaviour and
// checks if the Go template language works as intended.
func TestTextField_HTMLRendering(t *testing.T) {
	t.Parallel()

	t.Run("field building", func(t *testing.T) {
		t.Parallel()

		templ, err := template.New("").Parse(`
			<form>
				{{ .form.Name.Label }}
				{{ .form.Name.Input }}
				{{ .form.ImplicitField.Full }}
			</form>`)
		assert.NoError(t, err)

		buf := &bytes.Buffer{}
		err = templ.Execute(buf, map[string]any{
			"form": f.New(NameForm{
				Name: f.TextField("Your name"),
				// ImplicitField // not set on purpose to ensure rendering covers this case
			}),
		})
		assert.NoError(t, err)

		assert.Contains(t, buf.String(), "<label")
		assert.Contains(t, buf.String(), "<input")
		assert.Contains(t, buf.String(), "Your name")
		assert.Contains(t, buf.String(), "ImplicitField")
	})

	t.Run("field errors", func(t *testing.T) {
		t.Parallel()

		templ, err := template.New("").Parse(`
			<form>
				{{ .form.Name.Label }}
				{{ .form.Name.Input }}
				
				{{ if .form.Name.Errors }}inputElement has validation errors{{end}}
				{{ range .form.Name.Errors }}
					error: {{.}}
					error-key: {{.Key}}
					error-msg: {{.Message}}
				{{end}}
			</form>`)
		assert.NoError(t, err)

		form := f.New(NameForm{
			Name: f.TextField("Your name", f.Required()),
		})
		f.Validate(form, newRequest(""))

		buf := &bytes.Buffer{}
		err = templ.Execute(buf, map[string]any{
			"form": form,
		})
		assert.NoError(t, err)

		assert.Contains(t, buf.String(), "inputElement has validation errors")
		assert.Contains(t, buf.String(), "error: Your name is required")
		assert.Contains(t, buf.String(), "error-key: Your name")
		assert.Contains(t, buf.String(), "error-msg: is required")
	})

	t.Run("full", func(t *testing.T) {
		t.Parallel()

		templ, err := template.New("").Parse(`
			<form>
				{{ .form.Name.Full }}
			</form>`)
		assert.NoError(t, err)

		form := f.New(NameForm{
			Name: f.TextField("Your name", f.Required()),
		})
		f.Validate(form, newRequest(""))

		buf := &bytes.Buffer{}
		err = templ.Execute(buf, map[string]any{
			"form": form,
		})
		assert.NoError(t, err)

		assert.Contains(t, buf.String(), "Your name is required")
	})
}

func TestRequired(t *testing.T) {
	t.Parallel()

	tf := f.TextField("my-label", f.Required())
	assert.Contains(t, tf.Label(), " *")
	assert.Contains(t, tf.Input(), " required")
}

type NameForm struct {
	Name          f.Text
	ImplicitField f.Text
}
