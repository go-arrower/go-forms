package f

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// white box test. if it fails, feel free to delete it
func TestBase_SetValue(t *testing.T) {
	t.Parallel()

	type MyInvalidField struct {
		base
	}

	assert.PanicsWithValue(t, "go-forms: this field is implemented incorrectly: `base` assumes string type for value", func() {
		(&MyInvalidField{}).setValue(true)
	})
}
