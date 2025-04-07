package f_test

import (
	"testing"
	"time"

	"github.com/go-arrower/go-forms/f"
)

func TestDateTimeLocalField(t *testing.T) {
	t.Parallel()

	field := f.DateTimeLocalField("some",
		f.WithID(""),
		f.WithMax(time.Now()),
	)
	_ = field
}
