package f

type base struct {
	id           string
	label        string
	name         string
	value        string
	defaultValue string
}

var _ customField = (*base)(nil)

func (b *base) ID() string {
	return b.id
}

func (b *base) setID(id string) {
	b.id = id
}

func (b *base) Name() string {
	return b.name
}

func (b *base) setName(name string) {
	b.name = name
}

func (b *base) setValue(value any) {
	b.value = value.(string)
}
