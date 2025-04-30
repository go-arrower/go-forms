package f

type base struct {
	htmlID   string
	label    string
	htmlName string
	// value is of type string. This is a convenience as it works for many
	// input types. If an input type has a different type, it has to overwrite
	// this definition.
	value      string
	validators []func(field any, value string) error
	errors     []Error

	required     bool
	disabled     bool
	defaultValue string

	title     string
	form      string
	autofocus bool
}

var _ inputElement = (*base)(nil)

func (b *base) setBase(base base) {
	*b = base
}

func (b *base) id() string {
	return b.htmlID
}

func (b *base) setID(id string) {
	b.htmlID = id
}

func (b *base) name() string {
	return b.htmlName
}

func (b *base) setName(name string) {
	b.htmlName = name
}

func (b *base) setValue(value string) {
	b.value = value
}

func (b *base) validate() bool {
	panic("go-forms: this method MUST NOT be called on struct `base`, but on a field")
}
