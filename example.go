package main

import (
	"html/template"
	"net/http"

	"github.com/go-arrower/go-forms/f"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, _ *http.Request) {
		form := YourForm()

		templ, _ := template.New("").Parse(page)
		_ = templ.Execute(res, map[string]any{
			"form": form,
		})
	})

	_ = http.ListenAndServe(":8080", nil)
}

func YourForm() struct {
	FirstName f.Field
	LastName  f.Field
} {
	return struct {
		FirstName f.Field
		LastName  f.Field
	}{
		f.TextField("Your Firstname", f.Required()),
		f.TextField("Your Lastname"),
	}
}

const page = `<!DOCTYPE html>
<html>
<head>
    <style>
        form {
            width: 300px;
            margin: 10% auto 0;
        }
        .form-group {
            margin-bottom: 10px;
        }
        label {
            display: block;
            margin-bottom: 5px;
        }
    </style>
</head>
<body>
    <form>
        <div class="form-group">
            {{ .form.FirstName.Label }}
            {{ .form.FirstName.Input }}
        </div>
        <div class="form-group">
            {{ .form.LastName.Label }}
            {{ .form.LastName.Input }}
        </div>

		<input type="submit" value="Submit" />
    </form>
</body>
</html>`
