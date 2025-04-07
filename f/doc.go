// Package f implements form utility routines for easily working with forms.
//
// The focus is on quickly creating, validating, and rendering HTML forms.
// The package expects the developer to create valid forms.
// To keep the workflow quick and controller short,
// the API does not return any errors and panics instead.
//
// HTMLInputElement has a lot of possible attributes and in HTML it is easy
// to create invalid input elements. Browsers are quite forgiving,
// go-forms aims to nevertheless assist the developer to create as valid forms
// as possible.
// The options are typed according to their field to give compile-time
// feedback and prevent invalid combinations.
package f
