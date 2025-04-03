package f_test

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-arrower/go-forms/f"
)

func TestBooleanField(t *testing.T) {
	t.Run("defaults", func(t *testing.T) {
		t.Parallel()

		label := template.HTML(`<label for="label">Label</label>`)
		html := template.HTML(`<input type="text" id="label" name="label" value=""/>`)

		field := f.BooleanField("Label")
		assert.Equal(t, "label", field.ID())
		assert.Equal(t, "label", field.Name())
		assert.Equal(t, label, field.Label())
		assert.Equal(t, html, field.Input())
	})
}
