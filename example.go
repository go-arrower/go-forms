package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"

	"github.com/go-arrower/go-forms/f"
)

func main() {
	http.HandleFunc("/", formExample)

	_ = http.ListenAndServe(":8080", nil)
}

type MyForm struct {
	FirstName f.Text
	LastName  f.Text
	Pet       f.Text
}

func formExample(res http.ResponseWriter, req *http.Request) {
	form := f.New(MyForm{
		FirstName: f.TextField("my-name"),
	})

	if req.Method == http.MethodPost && f.Validate(form, req) {
		fn := form.FirstName.Value()
		fmt.Printf("submitted example 1: FirstName=%s\n", fn)
		fmt.Printf("submitted example 1: Pet=%s\n", form.Pet.Value())

		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	err := templ.Execute(res, map[string]any{
		"form": form,
	})
	fmt.Println(err)
}

var templ, _ = template.New("").Funcs(map[string]any{
	"old": func(form any) []any {
		var ret []any

		rf := reflect.ValueOf(form)
		fmt.Println(rf)
		fmt.Println(rf.Elem().Type())
		fmt.Println(rf.Elem().Kind())

		for i := range rf.Elem().NumField() {
			field := rf.Elem().Field(i)
			// check that the field implements f.field, to ensure it is a forms value, continue otherwise
			ret = append(ret, field.Interface())
		}

		return ret
	},
	"fields2": func(form any) []any {
		v := reflect.ValueOf(form).Elem()
		// if v.Kind() != reflect.Struct {
		// 	return nil
		// }

		fmt.Println("TMPL FUNC", v.NumField())
		out := make([]any, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			out[i] = v.Field(i).Addr().Interface()
			// out[i] = "HIER"
			fmt.Println("REFECT FIELD", i, v.Field(i).Type().Name(), v.Field(i).Kind())
		}

		return out
	},
}).Parse(page)

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
		<div class="form-group">
            {{ .form.Pet.Label }}
            {{ .form.Pet.Input }}
        </div>

		<hr />
		<hr />
		{{ range fields2 .form }}
			<div class="form-group">
				{{ .Full }}<-
			</div>
		{{ end }}

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
