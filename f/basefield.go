package f

type base struct {
	id           string
	label        string
	htmlName     string
	value        string
	defaultValue string
}

// var _ inputElement = (*base)(nil)

func (b *base) ID() string {
	return b.id
}

func (b *base) setID(id string) {
	b.id = id
}

func (b *base) name() string {
	return b.htmlName
}

func (b *base) setName(name string) {
	b.htmlName = name
}

func (b *base) setValue(value any) {
	b.value = value.(string)
}
