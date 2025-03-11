package f_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-arrower/go-forms/f"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		r        *http.Request
		f        any
		expValid bool
	}{
		"nil request":  {nil, struct{}{}, false},
		"nil form":     {newRequest(""), nil, false},
		"invalid form": {newRequest(""), invalidFormStruct{}, false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expValid, f.Validate(tt.r, tt.f))
		})
	}

	t.Run("form value populated", func(t *testing.T) {
		t.Parallel()

		form := struct {
			Name *f.Text // TODO make it work with f.Field
		}{f.TextField("param")}

		assert.True(t, f.Validate(newRequest("param=value"), form))
		assert.Equal(t, "value", form.Name.Value())
	})

	t.Run("form with all fields", func(t *testing.T) {
		t.Parallel()

		form := struct {
			Name    f.Field // *f.Text
			Checked f.Field // *f.Boolean
		}{
			Name: f.TextField("name-label",
				f.WithPlaceholder("Your lovely name"),
			),
			Checked: f.BooleanField("checked-label"),
		}

		assert.True(t, f.Validate(newRequest("param=value"), form))
	})
}

/*
	TODO for test cases
	* give http request to f.Validate
		* extract data for POST, PUT, PATCH requests
	* All the WithX methods
*/

func newRequest(data string) *http.Request {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req
}

type invalidFormStruct struct {
	Name string // field is not of type f.Field
}

func TestBuilder_Form(t *testing.T) {
	t.Parallel()

	formDef := f.Build().
		Text("Zuerst",
			f.WithID(""),
		).
		Checkbox("Als letztes")

	for field := range formDef.Fields2() {
		t.Log(field)
	}
}
