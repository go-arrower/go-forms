package f_test

import (
	"bytes"
	"context"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-arrower/go-forms/f"
)

var ctx = context.Background()

func TestTextField(t *testing.T) {
	t.Parallel()

	t.Run("defaults", func(t *testing.T) {
		t.Parallel()

		label := template.HTML(`<label for="label">Label</label>`)
		html := template.HTML(`<input type="text" id="label" name="label" value=""/>`)
		full := template.HTML(`<label for="label">Label</label><input type="text" id="label" name="label" value=""/>`)

		field := f.TextField("Label")

		assert.Equal(t, "label", field.ID())
		assert.Equal(t, "label", field.Name())
		assert.Equal(t, label, field.Label())
		assert.Equal(t, html, field.Input())
		assert.Equal(t, full, field.Full())

		assert.Empty(t, field.Value())
		assert.Nil(t, field.Errors())

		field.SetValue("my-val")
		assert.Equal(t, "my-val", field.Value())
	})

	t.Run("with id", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithID("my-id"))

		assert.Equal(t, "my-id", field.ID())
		assert.Contains(t, field.Input(), "my-id")
	})

	t.Run("with name", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithName("my-name"))

		assert.Equal(t, "my-name", field.Name())
		assert.Contains(t, field.Input(), "my-name")
	})

	t.Run("with value", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label", f.WithValue("my-value"))

		assert.Equal(t, "my-value", field.Value())
		assert.Contains(t, field.Input(), "my-value")
	})

	t.Run("with list", func(t *testing.T) {
		t.Parallel()

		field := f.TextField("Label",
			f.WithID("my-id"),
			f.WithList([]string{"A", "B"}),
		)

		assert.Contains(t, field.Input(), `<datalist id="my-id-datalist">`)
		assert.Contains(t, field.Input(), `<option value="A"></option>`)
		assert.Contains(t, field.Input(), `<option value="B"></option>`)
		assert.Contains(t, field.Input(), `</datalist>`)
	})
}

// This test case is to ensure the field works in its entirety.
// Rendering it through the arrower renderer ensures behaviour and
// checks if the Go template language works as intended.
func TestTextField_HTMLRendering(t *testing.T) {
	t.Parallel()

	t.Run("form building", func(t *testing.T) {
		t.Parallel()

		templ, err := template.New("").Parse(`
			<form>
				{{ .form.Name.Label }}
				{{ .form.Name.Input }}
			</form>`)
		assert.NoError(t, err)

		buf := &bytes.Buffer{}
		err = templ.Execute(buf, map[string]any{
			"form": NameForm{
				Name: f.TextField("Your name"),
			},
		})
		assert.NoError(t, err)

		assert.Contains(t, buf.String(), "<label")
		assert.Contains(t, buf.String(), "<input")
	})

	t.Run("form errors", func(t *testing.T) {
		t.Parallel()

		templ, err := template.New("").Parse(`
			<form>
				{{ .form.Name.Label }}
				{{ .form.Name.Input }}
				
				{{ if .form.Name.Errors }}Field has validation errors{{end}}
				{{ range .form.Name.Errors }}
					error: {{.}}
					error-key: {{.Key}}
					error-msg: {{.Message}}
				{{end}}
			</form>`)
		assert.NoError(t, err)

		form := NameForm{
			Name: f.TextField("Your name", f.Required()),
		}
		f.Validate(newRequest(""), form)

		buf := &bytes.Buffer{}
		err = templ.Execute(buf, map[string]any{
			"form": form,
		})
		assert.NoError(t, err)

		assert.Contains(t, buf.String(), "Field has validation errors")
		assert.Contains(t, buf.String(), "error: Your name is required")
		assert.Contains(t, buf.String(), "error-key: Your name")
		assert.Contains(t, buf.String(), "error-msg: is required")
	})
}

func TestRequired(t *testing.T) {
	t.Parallel()

	tf := f.TextField("my-label", f.Required())
	assert.Contains(t, tf.Label(), " *")
	assert.Contains(t, tf.Input(), " required")
}

type NameForm struct {
	Name f.Field
}
