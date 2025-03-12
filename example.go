package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-arrower/go-forms/f"
)

var templ, _ = template.New("").Parse(page)

func main() {
	http.HandleFunc("/", formIsStructExample)
	http.HandleFunc("/1", formIsStructExample)
	http.HandleFunc("/2", formBuilderExample)
	http.HandleFunc("/3", formBuilder2Example)

	_ = http.ListenAndServe(":8080", nil)
}

func formIsStructExample(res http.ResponseWriter, req *http.Request) {
	form := YourForm()

	if req.Method == http.MethodPost && f.Validate(req, form) {
		fn := form.FirstName.Value()
		fmt.Printf("submitted example 1: FirstName=%s\n", fn)

		http.Redirect(res, req, "/1", http.StatusSeeOther)
	}

	_ = templ.Execute(res, map[string]any{
		"form": form,
	})
}

func formBuilderExample(res http.ResponseWriter, req *http.Request) {
	form := f.Build().
		Text("FirstName", f.Required()). // the label has to match the key in the template
		Text("LastName")

	if req.Method == http.MethodPost { // && form.Validate(req) {
		fn := form.Fields["FirstName"] // .Value() // type any
		fmt.Printf("submitted example 2: FirstName=%s\n", fn)

		http.Redirect(res, req, "/2", http.StatusSeeOther)
	}

	err := templ.Execute(res, map[string]any{
		"form": form.Form(),
	})
	fmt.Println(err)
}

func formBuilder2Example(res http.ResponseWriter, req *http.Request) {
	form := f.New().
		Text("FirstName", f.Required()). // the label has to match the key in the template
		Text("LastName")

	if req.Method == http.MethodPost { // && form.Validate(req) {
		fn := (*form)["FirstName"] // .Value() // type any
		fmt.Printf("submitted example 2: FirstName=%s\n", fn)

		http.Redirect(res, req, "/3", http.StatusSeeOther)
	}

	err := templ.Execute(res, map[string]any{
		"form": form,
	})
	fmt.Println(err)
}

/*
	form := f.New(Method(f.POST)).Text("Your Name")
	err := templ.Execute(res, map[string]any{
		"form": form.View(),
	})

	// in the HTML
	{{ .form.Full }}
	{{ range .form  }}
		{{ .Full }}
	{{ end }}

	// Alternative: use template functions
	{{ full_form .form }}
*/

func YourForm() struct {
	FirstName *f.Text
	LastName  f.Field
} {
	return struct {
		FirstName *f.Text
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
			border: 1px solid silver;
			padding: 25px 35px;
        }
        .form-group {
            margin-bottom: 10px;
        }
        label {
            display: block;
            margin-bottom: 5px;
        }
		nav {
			width: 300px;
			margin: 10% auto 0;
			text-align: center;
		}
    </style>
</head>
<body>
    <form method="post">
        <div class="form-group">
            {{ .form.FirstName.Label }}
            {{ .form.FirstName.Input }}
			{{ range .form.FirstName.Errors }}
				<li>{{ . }}</li>
			{{ end }}
        </div>
        <div class="form-group">
            {{ .form.LastName.Label }}
            {{ .form.LastName.Input }}
        </div>

		<input type="submit" value="Submit" />
    </form>

	<nav>
		Other examples<br/>
		<a href="/1">Form Struct</a> - 
		<a href="/2">Form Builder I</a> - 
		<a href="/3">Form Builder II</a>
	</nav>
</body>
</html>`
